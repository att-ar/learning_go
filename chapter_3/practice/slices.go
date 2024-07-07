package main

import (
	"fmt"
	"sort"
)

func slices() {
	// not defining the number of items in the array, makes it a slice, which is a very mutable array
	var colors = []string{"Red","Green","Blue"} // []type{items}
	fmt.Println(colors)
	
	colors = append(colors, "Purple") // append is a built-in function to add items to a slice
	fmt.Println(colors)

	colors = colors[1:] // removing the first item
	fmt.Println(colors)

	colors = colors[:len(colors)-1] // removing the last item
	fmt.Println(colors)

	numbers := make([]int, 5) // make is a built-in function to create a slice with a length and capacity
	numbers[0] = 124
	numbers[1] = 1
	numbers[2] = 12
	numbers[3] = 12123
	numbers[4] = 90
	fmt.Println(numbers)

	numbers = append(numbers, 235) // adding a number to the slice
	fmt.Println(numbers)

	sort.Ints(numbers) // sorting the slice
	fmt.Println(numbers)
}