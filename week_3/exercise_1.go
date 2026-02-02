package main

import "fmt"

func main() {
	floatExample := 1.75
	// %T verb prints the type of the variable
	fmt.Printf("Working with a %T", floatExample)

	fmt.Println("\n***") // Added for spacing

	// %v verb prints the value of the variable
	fmt.Printf("Working with a %v", floatExample)

	fmt.Println("\n***") // Added for spacing

	yearsOfExp := 3
	reqYearsExp := 15
	// %d verb prints integer values
	fmt.Printf("I have %d years of Go experience and this job is asking for %d years.", yearsOfExp, reqYearsExp)

	fmt.Println("\n***") // Added for spacing

	stockPrice := 3.50
	// %.2f verb prints float values with 2 decimal places
	fmt.Printf("Each share of Gopher feed is $%.2f", stockPrice)
}
