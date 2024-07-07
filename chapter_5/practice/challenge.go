package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(calculate("1.11", "2", "+"))
	fmt.Println(calculate("1.11", "2", "-"))
	fmt.Println(calculate("1.11", "2", "*"))
}

func calculate(input1 string, input2 string, operation string) float64 {
    float1 := convertInputToValue(input1)
	float2 := convertInputToValue(input2)

	switch operation {
	case "+":
		return addValues(float1, float2)
	case "-":
		return subtractValues(float1, float2)
	case "*":
		return multiplyValues(float1, float2)
	case "/":
		return divideValues(float1, float2)
	default:
		fmt.Println("Invalid operation: " + operation)
	}
	return 0
}

func convertInputToValue(input string) float64 {
	float, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		panic("Invalid input: " + input)
	}
	return float
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
