package techpalace

import "strings"

// After ensuring it ran, we can consolidate and nest some functions.

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    starRepeat := strings.Repeat("*", numStarsPerLine)
    return starRepeat + "\n" + welcomeMsg + "\n" + starRepeat
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}
