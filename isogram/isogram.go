package isogram

import (
	"regexp"
	"strings"
)

func IsIsogram(word string) bool {
	r := regexp.MustCompile(`(\W)`)
	word = strings.ToLower(r.ReplaceAllString(word, ""))

	seen := make(map[rune]bool)
	for _, v := range word {
		if _, ok := seen[v]; ok {
			return false
		}

		seen[v] = true
	}

	return true
}
