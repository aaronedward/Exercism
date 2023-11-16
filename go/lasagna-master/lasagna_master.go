package lasagna

func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		return (2 * len(layers))
	} else {
		return (prepTime * len(layers))
	}
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (int, float64) {
	var noodleTotal = 0
	var sauceTotal = 0.0

	for i := 0; i < len(layers); i++ {
		if layers[i] == "sauce" {
			sauceTotal++
		} else if layers[i] == "noodles" {
			noodleTotal++
		} else {
			continue
		}
	}
	return noodleTotal * 50, sauceTotal * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient([]string, []string) {
	panic("")
}

// TODO: define the 'ScaleRecipe()' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
