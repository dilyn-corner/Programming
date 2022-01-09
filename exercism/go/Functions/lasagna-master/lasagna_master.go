package lasagna

func PreparationTime(layers []string, minutes int) int {
	if minutes == 0 {
		minutes = 2
	}
	return len(layers) * minutes
}

func Quantities(layers []string) (int, float64) {
	var qNoodle int
	var qSauce float64
	for _, ingredient := range layers {
		if ingredient == "noodles" {
			qNoodle += 50
		}
		if ingredient == "sauce" {
			qSauce += 0.2
		}
	}
	return qNoodle, qSauce
}

func AddSecretIngredient(friendsList, myList []string) {
	secretIngredient := friendsList[len(friendsList)-1]
	newList := myList
	newList[len(newList)-1] = secretIngredient
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	newQuantities := make([]float64, len(quantities))
	for i, ingredient := range quantities {
		newQuantities[i] = (ingredient / 2) * float64(portions)
	}
	return newQuantities
}
