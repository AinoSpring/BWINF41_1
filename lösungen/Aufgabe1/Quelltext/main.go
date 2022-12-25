package main

import "fmt"

func main() {
	var uncompletedSentences = ImportFiles()
	var book = ImportBook()
	for idx, uncompletedSentence := range uncompletedSentences {
		var sentenceChannel = make(chan string)
		fmt.Printf("Sentence possibilities for uncompleted sentence %v: \n", idx)
		go Search(uncompletedSentence, book, sentenceChannel)
		for sentence := range sentenceChannel {
			fmt.Printf("%#v\n", sentence)
		}
		fmt.Print("\n")
	}
}
