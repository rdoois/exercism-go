package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var res []string
	if len(rhyme) == 0 {
		return res
	}

	for i := 1; i < len(rhyme); i++ {
		phrase := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i])
		res = append(res, phrase)
	}

	res = append(res, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return res
}
