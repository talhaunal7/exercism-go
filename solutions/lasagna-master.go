//https://exercism.org/tracks/go/exercises/lasagna-master , functions
package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutes int) int {
	if minutes == 0 {
		return len(layers) * 2
	}
	return len(layers) * minutes
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		} else if v == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {

	myList[len(myList)-1] = friendsList[len(friendsList)-1]

}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaledQuantities []float64
	for _, v := range quantities {
		scaledQuantities = append(scaledQuantities, v*(float64(portions)/2))
	}
	return scaledQuantities
}
