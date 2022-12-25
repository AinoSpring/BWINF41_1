package main

import (
	"log"
	"os"
	"strings"
)

const (
	CONFIG = "config"
)

func ImportFiles() [][]Term {
	var termLists = make([][]Term, 0)
	data, err := os.ReadFile(CONFIG)
	if err != nil {
		log.Panic(err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSuffix(line, "\r")
		termListData, termListErr := os.ReadFile(line)
		if termListErr != nil {
			log.Panicf("%#v: %v", line, termListErr)
		}
		var termList = make([]Term, 0)
		for _, line := range strings.Split(strings.Replace(string(termListData), "\r", "", -1), "\n") {
			if line == "" {
				continue
			}
			termList = append(termList, Term{strings.Split(line, " ")[0], strings.Split(line, " ")[1]})
		}
		termLists = append(termLists, termList)
	}
	return termLists
}
