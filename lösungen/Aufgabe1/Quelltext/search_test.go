package main

import (
	"testing"
)

func sliceEqual[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestSearch(t *testing.T) {
	var text = "C# is Microsoft Java. Hello Hello World I am Aino ABC Python is bad"
	var sentence = "_ World _ am _"
	var result = []string{"Hello World I am Aino"}
	var answer = make([]string, 0)
	var ch = make(chan string)
	go Search(sentence, &text, ch)
	for possibleSentence := range ch {
		answer = append(answer, possibleSentence)
	}
	if !sliceEqual(result, answer) {
		t.Fatalf("Expected %#v, got %#v", result, answer)
	}
}
