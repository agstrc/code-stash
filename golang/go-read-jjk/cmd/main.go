package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

var outLogger = log.New(os.Stdout, "", log.LstdFlags)
var errLogger = log.New(os.Stdout, "", log.LstdFlags)
var conf = readConfig()

func main() {
	outLogger.Println("Checking for new JJK chapter at https://mangaplus.shueisha.co.jp/titles/100034")
	latestR := latestReleased()
	latestN, err := strconv.Atoi(conf["chapter"])

	// I could setup the config file to have a number type for chapter, but I'd still have to handle the case of
	// mismatched types. This model is easier as the JSON Unmarshaler checks for the correct types already and I
	// only have to make sure the "chapter" key is an actual int
	if err != nil {
		errLogger.Fatalln("Could not properly parse \"chapter\" key value:", err)
	}

	if latestR == latestN {
		outLogger.Println("No new chapter detected")
		os.Exit(0)
	}

	chapterStr := fmt.Sprint(latestR)
	msg := "Chapter " + chapterStr + " has been released!"
	outLogger.Println(msg)
	sendMessage(msg)
	conf["chapter"] = chapterStr

	cfgBytes, _ := json.MarshalIndent(conf, "", "  ")
	err = os.WriteFile(fileDir+confFile, cfgBytes, 0644)

	if err != nil {
		log.Fatalln(err)
	}
}
