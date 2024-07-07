package main

import (
	"fmt"
	"strconv"
	"strings"
)

// chapter 2 code challenge:

func calculate(value1 string, value2 string) float64 {
	// given two string parameters that can be parsed as float64
	float1, err1 := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	float2, err2 := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if ((err1 != nil) || (err2 != nil)) {
		fmt.Println(err1, err2)
		panic("Error parsing float64")
	}
	return float1 + float2
}

func main() {
	float := calculate("90.0 ", "9.8")
	fmt.Println(float)
}