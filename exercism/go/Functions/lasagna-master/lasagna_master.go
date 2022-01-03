package lasagna

func PreparationTime(layers []string, minutes int) int {
    var time int
    if minutes == 0 {
        time = len(layers) * 2
    } else {
        time = len(layers) * minutes
    }
    return time
}

func Quantities(layers []string) (int, float64) {
    qNoodle := 0
    qSauce := 0.0
    for i := 0; i < len(layers); i++ {
        switch layers[i] {
        case "noodles":
            qNoodle += 50
        case "sauce":
            qSauce +=0.2
        }
    }
        return qNoodle, qSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) []string {
    secretIngredient := friendsList[len(friendsList)-1]
    newList := append(myList, secretIngredient)
    return newList
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
    var newQuantities []float64
    var perPortion float64 = float64(portions) / 2.0
    for i := 0; i < len(quantities); i++ {
        newPortion := quantities[i] * perPortion
        newQuantities = append(newQuantities, newPortion)
    }
    return newQuantities
}
