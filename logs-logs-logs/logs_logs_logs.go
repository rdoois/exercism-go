package logs

import (
	"fmt"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
    logs := map[string]string{
        "U+2757": "recommendation",
        "U+1F50D": "search",
        "U+2600": "weather",
    }

    for _, c := range log {
        unicode := fmt.Sprintf("%U", c)
        value, exists := logs[unicode]
        if exists {
            return value
        }
    }

    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    var newLog string
    for _, char := range log {
        if char == oldRune {
            newLog = fmt.Sprintf("%s%c", newLog, newRune)        
            continue
        }

        newLog = fmt.Sprintf("%s%c", newLog, char)  
    }
    return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    var chars int
    for range log {
       chars++
    }
    return chars <= limit
}
