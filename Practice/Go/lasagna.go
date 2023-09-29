package lasagna

import "fmt"

const OvenTime int = 40

func main() {
	fmt.Println(RemainingOvenTime(30))
	fmt.Println(PreparationTime(2))
	fmt.Println(ElapsedTime(3, 20))
}

func RemainingOvenTime(actual int) int {

	return OvenTime - actual
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return (numberOfLayers*2 + actualMinutesInOven)
}
