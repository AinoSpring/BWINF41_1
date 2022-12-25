package main

import (
	"golang.org/x/exp/slices"
	"strings"
	"testing"
)

func TestGetVariableGroupsWith1LengthvowelGroups(t *testing.T) {
	var words = []string{"Hello", "World", "Golang", "Python", "Java", "Rust", "Haskell", "Erlang", "Fortran"}
	var solutions = []string{"e.o", "o", "o.a", "o", "a.a", "u", "a.e", "e.a", "o.a"}
	for idx, word := range words {
		var answer = GetVowelGroups(word)
		var solution = strings.Split(solutions[idx], ".")
		if len(solution) != len(answer) {
			t.Fatalf("On word %#v: exptected %v vowel groups, got %v: %#v", word, len(solution), len(answer), answer)
		}
		for idx2, vowelGroup := range answer {
			if strings.ToLower(solution[idx2]) != strings.ToLower(vowelGroup.value) {
				t.Fatalf("On word %#v: expected %v, got %v.", word, solution[idx2], vowelGroup.value)
			}
		}
	}
}

func TestGetVariableGroupsWithMoreThan1LengthVowelGroups(t *testing.T) {
	var words = []string{"Rain", "Meat", "Seat", "Pain", "Lie", "Paint", "Pie", "Queue", "Equal"}
	var solutions = []string{"ai", "ea", "ea", "ai", "ie", "ai", "ie", "ueue", "E.ua"}
	for idx, word := range words {
		var answer = GetVowelGroups(word)
		var solution = strings.Split(solutions[idx], ".")
		if len(solution) != len(answer) {
			t.Fatalf("On word %#v: exptected %v vowel groups, got %v: %#v", word, len(solution), len(answer), answer)
		}
		for idx2, vowelGroup := range answer {
			if strings.ToLower(solution[idx2]) != strings.ToLower(vowelGroup.value) {
				t.Fatalf("On word %#v: expected %v, got %v.", word, solution[idx2], vowelGroup.value)
			}
		}
	}
}

func TestGetEssentialVowelGroupWithOneVowelGroup(t *testing.T) {
	var words = []string{"Rain", "Meat", "Seat", "Pain", "Lie", "Paint", "Pie", "Queue", "Rust"}
	var solutions = []string{"ai", "ea", "ea", "ai", "ie", "ai", "ie", "ueue", "u"}
	for idx, word := range words {
		var answer = GetEssentialVowelGroup(word)
		if strings.ToLower(solutions[idx]) != strings.ToLower(answer.value) {
			t.Fatalf("On word %#v: expected %v, got %v.", word, solutions[idx], answer.value)
		}
	}
}

func TestGetEssentialVowelGroupWithMoreThanOneVowelGroup(t *testing.T) {
	var words = []string{"Hello", "Potato", "Golang", "Banana", "Java", "Equal", "Haskell", "Erlang", "Fortran"}
	var solutions = []string{"e", "a", "o", "a", "a", "e", "a", "e", "o"}
	for idx, word := range words {
		var answer = GetEssentialVowelGroup(word)
		if strings.ToLower(solutions[idx]) != strings.ToLower(answer.value) {
			t.Fatalf("On word %#v: expected %v, got %v.", word, solutions[idx], answer.value)
		}
	}
}

func TestCheckIfEndSame(t *testing.T) {
	var words = []string{"Potato.Tomato", "Python.Good", "Cat.Hat", "Hello.World", "Ball.Tall", "Yes.No", "Bridge.Fridge", "Good.Bad", "Cave.Gave"}
	var solutions = []bool{true, false, true, false, true, false, true, false, true}
	for idx, word := range words {
		var answer = CheckIfEndSame(strings.Split(word, ".")[0], strings.Split(word, ".")[1])
		if solutions[idx] != answer {
			t.Fatalf("On word %#v: expected %v, got %v.", word, solutions[idx], answer)
		}
	}
}

func TestCheckIfEssentialVowelGroupValid(t *testing.T) {
	var words = []string{"Potato", "No", "Fridge", "Chris", "Banana", "Tyler", "Right"}
	var solutions = []bool{true, true, true, false, true, false, true}
	for idx, word := range words {
		var answer = CheckIfEssentialVowelGroupValid(word)
		if solutions[idx] != answer {
			t.Fatalf("On word %#v: expected %v, got %v.", word, solutions[idx], answer)
		}
	}
}

func TestCheckIfWordsIncludeThem(t *testing.T) {
	var words = []string{"Grandmother.Mother", "Python.Good", "Motherboard.Board", "Hello.World", "Other.Otherwise", "Yes.No", "Another.Other", "Good.Bad"}
	var solutions = []bool{true, false, true, false, false, false, true, false}
	for idx, word := range words {
		var answer = CheckIfWordsIncludeThem(strings.Split(word, ".")[0], strings.Split(word, ".")[1])
		if solutions[idx] != answer {
			t.Fatalf("On word %#v: expected %v, got %v.", word, solutions[idx], answer)
		}
	}
}

func TestCheckIfRhyme(t *testing.T) {
	var words = []string{"Grandmother.Mother", "Fridge.Bridge", "Hello.World", "Potato.Tomato", "Yes.No", "Snow.Slow", "Good.Bad"}
	var solutions = []bool{false, true, false, true, false, true, false}
	for idx, word := range words {
		answer, rules := CheckIfRhyme(strings.Split(word, ".")[0], strings.Split(word, ".")[1])
		if solutions[idx] != answer {
			t.Fatalf("On word %#v: expected %v, got %v: %v", word, solutions[idx], answer, rules)
		}
	}
}

func TestCalculateRhymes(t *testing.T) {
	var words = []string{"Hello", "Save", "Fridge", "Cave", "Gave", "Bridge"}
	var solutions = []string{"cave.gave", "save.cave", "save.gave", "fridge.bridge"}
	var answer = CalculateRhymes(words)
	if len(answer) != len(solutions) {
		t.Fatalf("Expected length %v, got length %v.", len(answer), len(solutions))
	}
	for _, rhyme := range answer {
		if !(slices.Contains(solutions, strings.ToLower(rhyme.word1)+"."+strings.ToLower(rhyme.word2)) || slices.Contains(solutions, strings.ToLower(rhyme.word2)+"."+strings.ToLower(rhyme.word1))) {
			t.Fatalf("Got rhyme %v", rhyme)
		}
	}
}
