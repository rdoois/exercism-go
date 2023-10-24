package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
    res := make(map[string]int)
    for k, v := range in {
        for _, c := range v {
            res[strings.ToLower(c)] = k
        }
    }

    return res
}
