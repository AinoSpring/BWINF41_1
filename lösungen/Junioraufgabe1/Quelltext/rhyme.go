package main

import (
	"strings"

	"golang.org/x/exp/slices"
)

var (
	VOWELS = []string{"a", "e", "i", "o", "u", "ä", "ö", "ü"}
)

type VowelGroup struct {
	value string
	idx   uint
}

type Rhyme struct {
	word1 string
	word2 string
}

func GetVowelGroups(word string) (vowelGroups []VowelGroup) {
	word += " "
	vowelGroups = make([]VowelGroup, 0)
	var currentVowelGroup VowelGroup = VowelGroup{value: ""}
	for idx, character := range strings.Split(word, "") {
		if slices.Contains(VOWELS, strings.ToLower(character)) {
			if "" == currentVowelGroup.value {
				currentVowelGroup = VowelGroup{
					value: character,
					idx:   uint(idx),
				}
			} else {
				currentVowelGroup.value += character
			}
		} else {
			if "" == currentVowelGroup.value {
				continue
			}
			vowelGroups = append(vowelGroups, currentVowelGroup)
			currentVowelGroup = VowelGroup{value: ""}
		}
	}
	return
}

func GetEssentialVowelGroup(word string) VowelGroup {
	var vowelGroups = GetVowelGroups(word)
	if len(vowelGroups) > 1 {
		return vowelGroups[len(vowelGroups)-2]
	}
	if len(vowelGroups) < 1 {
		return VowelGroup{value: ""}
	}
	return vowelGroups[0]
}

func CheckIfEndSame(word1 string, word2 string) bool {
	var word1EssentialVowelGroup = GetEssentialVowelGroup(word1)
	var word2EssentialVowelGroup = GetEssentialVowelGroup(word2)
	return strings.Join(strings.Split(word1, "")[word1EssentialVowelGroup.idx:], "") == strings.Join(strings.Split(word2, "")[word2EssentialVowelGroup.idx:], "")
}

func CheckIfEssentialVowelGroupValid(word string) bool {
	var essentialVowelGroup = GetEssentialVowelGroup(word)
	if "" == essentialVowelGroup.value {
		return false
	}
	return float32(len(strings.Split(word, "")[essentialVowelGroup.idx:])) >= ((float32(len(strings.Split(word, "")))) / 2)
}

func CheckIfWordsIncludeThem(word1 string, word2 string) bool {
	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)
	var word1Length = len(strings.Split(word1, ""))
	var word2Length = len(strings.Split(word2, ""))
	if word1Length >= word2Length {
		return word1[word1Length-word2Length:] == word2
	} else if word1Length < word2Length {
		return word1 == word2[word2Length-word1Length:]
	} else {
		return false
	}
}

func CheckIfRhyme(word1 string, word2 string) (bool, []uint8) {
	word1 = strings.TrimSuffix(word1, " ")
	word2 = strings.TrimSuffix(word2, " ")
	var rhyme = true
	var rules = make([]uint8, 0)
	for _, word := range []string{word1, word2} {
		if !CheckIfEssentialVowelGroupValid(word) {
			rhyme = false
			rules = append(rules, 2)
		}
	}
	if rhyme {
		if !CheckIfEndSame(word1, word2) {
			rhyme = false
			rules = append(rules, 1)
		}
	}
	if CheckIfWordsIncludeThem(word1, word2) {
		rhyme = false
		rules = append(rules, 3)
	}
	return rhyme, rules
}

func CalculateRhymes(words []string) []Rhyme {
	var rhymes = make([]Rhyme, 0)
	for idx, word := range words {
		for i := idx + 1; i < len(words); i++ {
			if ok, _ := CheckIfRhyme(word, words[i]); ok {
				rhymes = append(rhymes, Rhyme{word1: word, word2: words[i]})
			}
		}
	}
	return rhymes
}
