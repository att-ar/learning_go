package main

import (
	"fmt"
	"math"
	"math/rand"
)

func convertToMap(items []string) map[string]float64 {
	// Create a map object
	result := make(map[string]float64)
    // Your code goes here
	for _, fruit := range items {
		result[fruit] = math.Round(rand.Float64() * 100)
	}

    return result
}

func main () {
	items := []string{"apple", "banana", "cherry"}
	result := convertToMap(items)
	fmt.Println(result)
	fmt.Printf("Type %T\n", result)

}