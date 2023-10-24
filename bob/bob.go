// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
    remark = strings.Trim(remark, " \t\r\n")
    if isSilence(remark) {
        return "Fine. Be that way!"
    } 
    if isYelling(remark) && isQuestion(remark) {
        return "Calm down, I know what I'm doing!"
    }
    if isYelling(remark) {
        return "Whoa, chill out!"
    }
    if isQuestion(remark) {
        return "Sure."
    }
    return "Whatever."
}

func isSilence(remark string) bool {
    return remark == ""   
}

func isYelling(remark string) bool {
    alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    return strings.ToUpper(remark) == remark && strings.ContainsAny(remark, alphabet)
}

func isQuestion(remark string) bool {
    return strings.HasSuffix(remark, "?")
}
