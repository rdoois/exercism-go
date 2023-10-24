package strain

func Keep[T any](input []T, predicate func(item T) bool) []T {
    var res []T
    for _, item := range input {
        if predicate(item) {
            res = append(res, item)
        }
    }
    return res
}

func Discard[T any](input []T, predicate func(item T) bool) []T {
    var res []T
    for _, item := range input {
        if !predicate(item) {
            res = append(res, item)
        }
    }
    return res
}

