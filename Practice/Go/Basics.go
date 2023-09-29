package main

import "fmt"

func main() {


	fmt.Println(RemainingOvenTime(30))
    
}


func RemainingOvenTime(actual int) int {
	const OvenTime int = 40

	return OvenTime - actual
}

func PreparationTime(numberOfLayers int) int {
    return numberOfLayers + 2 
}

