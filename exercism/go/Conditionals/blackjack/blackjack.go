package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch card {
    case "two":
        return 2
    case "three":
        return 3
    case "four":
        return 4
    case "five":
        return 5
    case "six":
        return 6
    case "seven":
        return 7
    case "eight":
        return 8
    case "nine":
        return 9
    case "ten", "jack", "queen", "king":
        return 10
    case "ace":
        return 11
    default:
        return 0
    }
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
    x := ParseCard(card1)
    y := ParseCard(card2)
    value := x + y
    switch value {
    case 21:
        return true
    default:
        return false
    }
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
    switch isBlackjack {
    case false:
        return "P"
    case true:
        if dealerScore < 10 {
            return "W"
        } else {
            return "S"
        }
    default:
        return "S"
    }
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
    switch {
    case 16 < handScore:
        return "S"
    case handScore < 12:
        return "H"
    default:
        if dealerScore > 6 {
            return "H"
        } else {
            return "S"
        }
    }
}
