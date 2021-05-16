package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Chapter int    `json:"chapter"`
	Webhook string `json:"webhook"`
}

var outLogger = log.New(os.Stdout, "", log.LstdFlags)
var errLogger = log.New(os.Stdout, "", log.LstdFlags)
var conf = readConfig()

func main() {
	outLogger.Println("Checking for new JJK chapter at https://mangaplus.shueisha.co.jp/titles/100034")
	latestR := latestReleased()
	latestN := conf.Chapter

	if latestR == latestN {
		outLogger.Println("No new chapter detected")
		os.Exit(0)
	}

	chapterStr := fmt.Sprint(latestR)
	msg := "Chapter " + chapterStr + " has been released!"
	outLogger.Println(msg)
	sendMessage(msg)
	conf.Chapter = latestR

	cfgBytes, _ := json.MarshalIndent(conf, "", "  ")
	err := os.WriteFile(fileDir+confFile, cfgBytes, 0644)

	if err != nil {
		log.Fatalln(err)
	}
}
