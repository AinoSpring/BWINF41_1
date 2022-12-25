package main

import (
	"strings"
)

func CheckWords(word string, wordInBook string) bool {
	word = strings.ToLower(word)
	wordInBook = strings.ToLower(wordInBook)
	if "_" == word {
		return true
	}
	return word == wordInBook
}

func Search(sentence string, text *string, returnChannel chan string) {
	defer close(returnChannel)
	var sentenceSlice = strings.Split(sentence, " ")
	var textSlice = strings.Split(*text, " ")
	for idx, word := range textSlice {
		if CheckWords(sentenceSlice[0], word) {
			var isSame = true
			for i := 1; i < len(sentenceSlice); i++ {
				if idx+i >= (len(textSlice) - 1) {
					isSame = false
				} else {
					if !CheckWords(sentenceSlice[i], textSlice[idx+i]) {
						isSame = false
						break
					}
				}
			}
			if isSame {
				returnChannel <- strings.Join(textSlice[idx:idx+len(sentenceSlice)], " ")
			}
		}
	}
}
