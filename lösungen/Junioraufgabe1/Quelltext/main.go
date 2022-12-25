package main

import "fmt"

func main() {
	var wordLists = importFiles()
	for idx, wordList := range wordLists {
		var rhymes = CalculateRhymes(wordList)
		fmt.Printf("Rhymes from wordlist %v: \n", idx)
		for _, rhyme := range rhymes {
			fmt.Printf("%#v, %#v\n", rhyme.word1, rhyme.word2)
		}
		fmt.Print("\n")
	}
}
