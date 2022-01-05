package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
    runeMap := map[rune]string {
        '❗': "recommendation",
        '🔍': "search",
        '☀' : "weather",
    }

    for _, char := range log {
        if _, r := runeMap[char]; r {
            return runeMap[char]
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
    var logLine string
    for _, char := range log {
        if char == oldRune {
            char = newRune
        }
        logLine += string(char)
    }
    return logLine
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    charCount := utf8.RuneCountInString(log)
    return charCount <= limit
}
