package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}

	var sum int
	use := false
	for i := len(id) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}
		if use {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		use = !use
	}

	return sum%10 == 0
}
