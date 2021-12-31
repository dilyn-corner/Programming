package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    floatRate := float64(productionRate)
    return (floatRate * successRate)/100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    floatProductionRate := float64(productionRate)
    intPerMinute := int(((floatProductionRate) * (successRate / 100)) / 60)
    return intPerMinute
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
    /*
    Presumably we can simply do our arithmetic and THEN convert and return
    that value without creating a bunch of variables, doing arithmetic on
    them, converting, and then returning.
    Theoretically, this might save time...
    */
    return uint((95000 * (carsCount / 10)) + (10000 * (carsCount % 10)))
}
