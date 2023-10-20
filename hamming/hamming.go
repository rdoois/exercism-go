package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("")
	}

	var dif int
	for i := range a {
		if a[i] != b[i] {
			dif++
		}
	}

	return dif, nil
}
