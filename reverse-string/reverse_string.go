package reverse

func Reverse(input string) string {
    var res string
    for _, c := range input {
        res = string(c) + res
    }
    return res
}
