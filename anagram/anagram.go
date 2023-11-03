package anagram

import (
	"fmt"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var res []string
	for _, candidate := range candidates {
		if compare(subject, candidate) {
			res = append(res, candidate)
		}
	}
	return res
}

func compare(subject, candidate string) bool {
	if subject == "good" {
		fmt.Print("")
	}
	if strings.EqualFold(subject, candidate) {
		return false
	}
	var word1, word2 []string
	if len(subject) > len(candidate) {
		word1 = strings.Split(strings.ToLower(subject), "")
		word2 = strings.Split(strings.ToLower(candidate), "")
	} else {
		word1 = strings.Split(strings.ToLower(candidate), "")
		word2 = strings.Split(strings.ToLower(subject), "")
	}

	for _, a := range word1 {
		hasLetter := false

		for i, b := range word2 {
			if a == b {
				hasLetter = true
				word2 = remove(word2, i)
				break
			}
		}
		if !hasLetter {
			return false
		}
	}
	return true
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
