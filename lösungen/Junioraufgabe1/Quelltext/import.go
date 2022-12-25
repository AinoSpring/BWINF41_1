package main

import (
	"log"
	"os"
	"strings"
)

const (
	CONFIG = "config"
)

func importFiles() [][]string {
	var wordLists = make([][]string, 0)
	data, err := os.ReadFile(CONFIG)
	if err != nil {
		log.Panic(err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSuffix(line, "\r")
		wordListData, wordListErr := os.ReadFile(line)
		if wordListErr != nil {
			log.Panicf("%#v: %v", line, wordListErr)
		}
		wordLists = append(wordLists, strings.Split(strings.Replace(string(wordListData), "\r", "", -1), "\n"))
	}
	return wordLists
}
