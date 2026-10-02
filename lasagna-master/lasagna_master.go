package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, meanPrepTime int) int {
	if meanPrepTime == 0 {
		meanPrepTime = 2
	}
	return len(layers) * meanPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsIngredients, myIngredients []string) {
	myIngredients[len(myIngredients)-1] = friendsIngredients[len(friendsIngredients)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	newQuantities := make([]float64, len(quantities))
	for i, quantity := range quantities {
		newQuantities[i] = quantity * float64(portions) / 2
	}
	return newQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
