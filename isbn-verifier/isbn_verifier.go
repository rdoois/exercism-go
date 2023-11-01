package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	size := len(isbn)
	if size != 10 {
		return false
	}

	var res int
	for k, v := range isbn {
		n, err := strconv.Atoi(string(v))
		if err != nil {
			if k != (size-1) || v != 'X' {
				return false
			}

			if err != nil {
				n = 10
			}
		}

		res += n * (size - k)
	}

	return res%11 == 0
}
