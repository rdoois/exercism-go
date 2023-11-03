package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 || span > len(digits) || len(digits) < 1 {
		return 0, errors.New("")
	}
	var res int64
	for i := 0; i <= len(digits)-span; i++ {
		serie := digits[i : i+span]
		mult := int64(1)
		for _, v := range serie {
			d, err := strconv.ParseInt(string(v), 10, 64)
			if err != nil {
				return 0, errors.New("")
			}
			mult = mult * d
		}
		if mult > res {
			res = mult
		}
	}
	return res, nil
}
