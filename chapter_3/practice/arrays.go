package main

import (
	"fmt"
)

func arrays() {
	var colors [3]string // [number of items]type of items
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"
	fmt.Println(colors)
	fmt.Println("Arrays")

	var numbers = [5]int{5,6,3,2,5} // [number of items]type of items{items}
	fmt.Println(numbers)

	fmt.Println("Number of colors:", len(colors))
	fmt.Println("Number of numbers:", len(numbers))
}