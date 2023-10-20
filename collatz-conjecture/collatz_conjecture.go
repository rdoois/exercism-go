package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("")
	}

	var steps int
	for steps = 0; n > 1; steps++ {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}

	return steps, nil

}
