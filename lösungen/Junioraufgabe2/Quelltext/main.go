package main

import "fmt"

func main() {
	var termLists = ImportFiles()
	for idx, termList := range termLists {
		var greatest = Greatest(Sort(termList))
		fmt.Printf("\nGreatest container from termlist %v: \n", idx)
		if greatest == "-1" {
			fmt.Printf("Greatest container could not be determined.")
		} else {
			fmt.Printf("%#v", greatest)
		}
		fmt.Print("\n")
	}
}
