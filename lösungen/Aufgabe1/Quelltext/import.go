package main

import (
	"log"
	"os"
	"strings"
)

const (
	CONFIG = "config"
)

var (
	IGNORED_CHARACTERS = []string{"»", "«", ",", ".", "?", "!", "_", "\""}
)

func ImportBook() *string {
	data, err := os.ReadFile(CONFIG)
	if err != nil {
		log.Panic(err)
	}
	bookData, bookErr := os.ReadFile(strings.TrimSuffix(strings.Split(string(data), "\n")[0], "\r"))
	if bookErr != nil {
		log.Panic(bookErr)
	}
	var bookString = string(bookData)
	bookString = strings.Replace(bookString, "\n", " ", -1)
	bookString = strings.Replace(bookString, "\r", "", -1)
	bookString = strings.Replace(bookString, "  ", " ", -1)
	for _, ignoredCharacter := range IGNORED_CHARACTERS {
		bookString = strings.Replace(bookString, ignoredCharacter, "", -1)
	}
	/*
		// For debugging purposes

		fo, err := os.Create("formattedBook.txt")
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()
		_, fErr := fmt.Fprintf(fo, "%v", bookString)
		if fErr != nil {
			return nil
		}
	*/
	return &bookString
}

func ImportFiles() []string {
	var sentences = make([]string, 0)
	data, err := os.ReadFile(CONFIG)
	if err != nil {
		log.Panic(err)
	}
	for _, line := range strings.Split(string(data), "\n")[1:] {
		line = strings.TrimSuffix(line, "\r")
		sentenceData, sentenceErr := os.ReadFile(line)
		if sentenceErr != nil {
			log.Panicf("%#v: %v", line, sentenceErr)
		}
		sentences = append(sentences, strings.Replace(strings.Replace(string(sentenceData), "\r", "", -1), "\n", "", -1))
	}
	return sentences
}
