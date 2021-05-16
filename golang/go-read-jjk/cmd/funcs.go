package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

const confFile = "config.json"

var fileDir = exeDir()

// exeDir returns an absolute path to the directory of the currently running executable
func exeDir() string {
	path, err := os.Executable()
	if err != nil {
		errLogger.Fatalln(err)
	}
	return filepath.Dir(path) + "/"
}

// readConfig returns a map which represents the configuration file
func readConfig() Config {
	bytes, err := os.ReadFile(fileDir + confFile)
	if err != nil {
		errLogger.Println("Failed to read configuration failed:", err)
		invalidConfigFile()
	}

	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		errLogger.Println("Failed to parse configuration file:", err)
		invalidConfigFile()
	}

	_, err = url.Parse(config.Webhook)
	if err != nil {
		errLogger.Println("Failed to parse discord webhook URL:", err)
		os.Exit(1)
	}

	return config
}

// invalidConfigFile writes the default configuration file into the appropriate location and exits the program
func invalidConfigFile() {
	configBytes, _ := json.MarshalIndent(
		map[string]interface{}{
			"chapter": 1,
			"webhook": "",
		},
		"", "  ")

	err := os.WriteFile(fileDir+confFile, configBytes, 0644)
	if err != nil {
		log.Fatalln("Could not create configurantion file:", err)
	}
	errLogger.Fatalf("Configuration file created at %s. Write your Discorrd webhook URL into it.", fileDir+confFile)
}

// latestReleased returns the number of the latest released JJK chapter
func latestReleased() int {
	var chaptersDiv string
	ctx, cancel := chromedp.NewContext(context.Background())

	// browser allocation. Does not need timeout
	if err := chromedp.Run(ctx); err != nil {
		log.Fatal(err)
	}
	cancel()

	// 1 minute timeout while waiting for the chapters to show up
	ctx, cancel = context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// since, at the time of writing, I am not sure if the reference website uses a consistente HTML layout, I mostly
	// played it safe. So the selection selects complete div which lists all available chapters. With that, I make use
	// of some simple regex to extract all the chapters that are listed.

	ctx, _ = chromedp.NewContext(ctx)
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://mangaplus.shueisha.co.jp/titles/100034"),
		chromedp.Text("#app > div:nth-child(2) > div.styles-module_container_1rtol > "+
			"div.styles-module_mainContainer_2tQWW > div > div > div.TitleDetail-module_flexContainer_1oGb4 > main > "+
			"div:nth-child(1)", &chaptersDiv),
	); err != nil {
		log.Fatal(err)
	}

	exp := regexp.MustCompile("#[0-9]+")
	chapters := exp.FindAllString(chaptersDiv, 16)

	if len(chapters) == 0 {
		log.Fatal("No chapters found")
	}

	// gets the latest chapter and removes the leading #
	latestStr := chapters[len(chapters)-1][1:]
	latestInt, err := strconv.Atoi(latestStr)
	if err != nil {
		msg := fmt.Sprintf("Failed to conver matched chapter regexp \"%s\" to int", latestStr)
		sendMessage(msg)
		errLogger.Fatalln(msg)
	}

	return latestInt
}

// sendMessage send a Discord message to the webhook specified in the configuration map
func sendMessage(msg string) {
	bodyBytes, _ := json.Marshal(map[string]string{
		"content": msg,
	})
	bodyBuffer := bytes.NewBuffer(bodyBytes)

	// 5 attempts to send the webhook message
	for i := 0; i < 5; i++ {
		resp, err := http.Post(conf.Webhook, "application/json", bodyBuffer)
		if err != nil {
			log.Println("Failed to send Discord webhook message:", err)
			time.Sleep(time.Second)
			continue
		} else if resp.StatusCode > 200 && resp.StatusCode < 299 {
			return
		}
	}

	log.Fatalln("Could not send Discord webhook message")
}
