package main

import (
	"fmt"
)

func structs() {
	poodle := Dog{"Poodle", 10} // constructors aren't methods, it is object definition so it needs curly braces
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle) // %+v prints the field names
	fmt.Println("Structs")
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	poodle.Weight = 5
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
}

// exported/public struct:

// Dog is a struct
type Dog struct {
	// exported/public fields
	Breed string
	Weight int
	// private fields
	// lowercase
}