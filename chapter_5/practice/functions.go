package main

import (
	"fmt"
)


func functions() {
	fmt.Println("Functions")
	doSomething()
	fmt.Println(addValues(1, 2))
	fmt.Println(addAllValues(1, 2, 3, 4, 5))
}

func doSomething() {
	fmt.Println("Doing something")
}

// since value1 and value2 are of the same type, we can write it like this (int only once after value2)
func _addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, len(values)
}