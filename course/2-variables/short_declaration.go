package main

import "fmt"

func main() {
	
	/* Go follows the printf tradition from the C language.
    fmt.Printf - Prints a formatted string to standard out.
    fmt.Sprintf() - Returns the formatted string */

	// var empty string
	empty := ""
	fmt.Printf("This is an empty string: %s\n", empty)
	numCars := 10 // inferred to be an integer
	fmt.Printf("He owns %d cars\n", numCars)

	temperature := 0.0 // temperature is inferred to be a floating point value because it has a decimal point
	fmt.Println(fmt.Sprintf("I am %.2f", temperature))
	
	var isFunny = true // isFunny is inferred to be a boolean
	fmt.Printf("The show was funny: %v\n", isFunny)

	// Multiple assignment
	mileage, company := 80276, "BMW"
	fmt.Println(mileage, company)
}