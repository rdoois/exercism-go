package romannumerals

import (
	"errors"
	"strconv"
)

var unit = map[int]string{
	1: "I",
	2: "II",
	3: "III",
	4: "IV",
	5: "V",
	6: "VI",
	7: "VII",
	8: "VIII",
	9: "IX",
}

var ten = map[int]string{
	1: "X",
	2: "XX",
	3: "XXX",
	4: "XL",
	5: "L",
	6: "LX",
	7: "LXX",
	8: "LXXX",
	9: "XC",
}

var hundred = map[int]string{
	1: "C",
	2: "CC",
	3: "CCC",
	4: "CD",
	5: "D",
	6: "DC",
	7: "DCC",
	8: "DCCC",
	9: "CM",
}

var thousand = map[int]string{
	1: "M",
	2: "MM",
	3: "MMM",
}

var numbers = map[int]map[int]string{
	0: unit,
	1: ten,
	2: hundred,
	3: thousand,
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("")
	}

	var res string
	in := strconv.Itoa(input)
	size := len(strconv.Itoa(input)) - 1

	for i := size; i >= 0; i-- {
		kind := numbers[size-i]
		n, _ := strconv.Atoi(string(in[i]))
		res = kind[n] + res
	}

	return res, nil
}
