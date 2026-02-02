package main

import "fmt"

func main() {
	fmt.Println("What would you like for lunch?")

	// Declare food variable below:
	var food string

	// Scan for user input below:
	fmt.Scan(&food)

	fmt.Printf("Sure, we can have %v for lunch.", food)
}
