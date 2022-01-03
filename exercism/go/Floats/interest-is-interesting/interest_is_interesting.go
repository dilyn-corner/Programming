package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    var interestRate float32 = 0
    if balance < 0 {
        interestRate = 3.213
    } else if 0 <= balance && balance < 1000 {
        interestRate = .5
    } else if 1000 <= balance && balance < 5000 {
        interestRate = 1.621
    } else if balance >= 5000 {
        interestRate = 2.475
    }
    return interestRate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    var interestAmount float64 = 0
    interestAmount = balance * float64(InterestRate(balance)) / 100
    return interestAmount
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
    balance += Interest(balance)
    return balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    var x int = 0
    for i := balance; i < targetBalance; i = i + Interest(i) {
        x++
    }
    return x
}
