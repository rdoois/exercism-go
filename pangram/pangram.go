package pangram

import "strings"

func IsPangram(input string) bool {
    input = strings.ToLower(input)
    for c := 'a'; c <= 'z'; c++ {
        if strings.IndexRune(input, c) == -1 {
            return false
        }
    }

    return true
}
