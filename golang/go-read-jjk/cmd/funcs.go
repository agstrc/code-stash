package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
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
func readConfig() map[string]string {
	bytes, err := os.ReadFile(fileDir + confFile)
	if err != nil {
		errLogger.Println("Failed to read configuration failed:", err)
		invalidConfigFile()
	}

	var configMap map[string]string
	err = json.Unmarshal(bytes, &configMap)
	if err != nil {
		errLogger.Println("Failed to parse configuration file:", err)
		invalidConfigFile()
	}

	// looks for the required keys on the configuration file map
	for _, str := range []string{"chapter", "webhook"} {
		if key, exists := configMap[str]; !exists {
			errLogger.Printf("Configuration file lacks the \"%s\" key", key)
			invalidConfigFile()
		}
	}

	return configMap
}

// invalidConfigFile writes the default configuration file into the appropriate location and exits the program
func invalidConfigFile() {
	configBytes, _ := json.MarshalIndent(
		map[string]string{
			"chapter": "1",
			"webhook": "",
		},
		"", "  ")

	err := os.WriteFile(fileDir+confFile, configBytes, 0644)
	if err != nil {
		log.Fatalln("Could not create configurantion file:", err)
	}
	errLogger.Fatalf("Configuration file created at %s. Write your Discorrd webhook URL into it.", fileDir+confFile)
}

// latestReleased returns the most recent JJK chapter listed on the website
func latestReleased() int {
	resp, err := http.Get("https://jujutsukaisen.online/")
	if err != nil {
		errLogger.Fatalln(err)
	}
	defer resp.Body.Close()

	html, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sel := html.Find("li.row:nth-child(1)").First()
	if sel.Length() == 0 {
		msg := "Failed to find latest chapter on https://jujutsukaisen.online/. No match found for the given CSS" +
			"selector"
		sendMessage(msg)
		errLogger.Fatalln(msg)
	}

	// no error handling needed. The expression is known at compile time
	exp, _ := regexp.Compile("[0-9]+")
	chapterStr := exp.FindString(sel.Text())

	chapterInt, err := strconv.Atoi(chapterStr)
	if err != nil {
		msg := fmt.Sprintf("Failed to conver matched chapter regexp \"%s\" to int", chapterStr)
		sendMessage(msg)
		errLogger.Fatalln(msg)
	}

	// at time of writing, the website always has a "preview" link to a unreleased chapter
	return chapterInt - 1
}

// sendMessage send a Discord message to the webhook specified in the configuration map
func sendMessage(msg string) {
	bodyBytes, _ := json.Marshal(map[string]string{
		"content": msg,
	})
	bodyBuffer := bytes.NewBuffer(bodyBytes)

	// 5 attempts to send the webhook message
	for i := 0; i < 5; i++ {
		resp, err := http.Post(conf["webhook"], "application/json", bodyBuffer)
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
