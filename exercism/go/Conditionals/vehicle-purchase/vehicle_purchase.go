package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    if kind == "car" || kind == "truck" {
        return true
    } else {
        return false
    }
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
    if option1 < option2 {
        return option1 + " is clearly the better choice."
    } else {
        return option2 + " is clearly the better choice."
    }
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    if age <= 3 {
        return float64(originalPrice * 4/5)
    } else if age >= 10 {
        return float64(originalPrice * 1/2)
    } else {
        return float64(originalPrice * 7/10)
    }
}
