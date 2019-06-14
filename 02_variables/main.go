package main

import "fmt"

func main() {
	// constant
	const country string = "Bangladesh"

	// multiple variable at once
	var firstName, lastName = "Mohammad", "Rezve"

	// with type
	var state string = "Dhaka"

	//short hand
	age := 12

	fmt.Printf("%s %s, age %d from %s, %s\n", firstName, lastName, age, state, country)
}
