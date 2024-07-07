package main

import (
	"fmt"
)

func for_loops() {
	colors := []string{"red", "green", "blue"}
	fmt.Println("Colors:", colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
	
	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println(color)
	}

	value := 1
	for value < 10 {
		fmt.Println("Value:", value)
		value*=2
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:",sum)
		if sum > 200 {
			goto theEnd
		}
		// can also use break and continue
	}
	theEnd : fmt.Println("End of Program") // this is a label, and it is used with goto
}