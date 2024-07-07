package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func challenge() {
	jsonString := `[{"name":"apple","price":2.99,"quantity":3},
	{"name":"orange","price":1.50,"quantity":8},
	{"name":"banana","price":0.49,"quantity":12}]`

	result := getCartFromJson(jsonString)
	for _, item := range result {
		fmt.Printf("%+v\n", item)
	}

	fmt.Println("\nUsing decoder.More():")
	resultNew := getCartFromJsonWithMore(jsonString)
	for _, item := range resultNew {
		fmt.Printf("%+v\n", item)
	}
}

type cartItem struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// Given the context of your challenge.go file,
// let's say you want to modify getCartFromJson to handle a stream of JSON objects that might be separated by newlines or other delimiters, rather than a single JSON array.
// You could use decoder.More() to continuously decode objects until there are no more available.

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []cartItem {
	var cart []cartItem
	decoder := json.NewDecoder(strings.NewReader(jsonString))
	err := decoder.Decode(&cart)
	if err != nil {
		panic(err)
	}
	return cart
}

// Modified getCartFromJson to use decoder.More()
func getCartFromJsonWithMore(jsonString string) []cartItem {
	var cart []cartItem
	decoder := json.NewDecoder(strings.NewReader(jsonString))

	// Loop while there are more JSON objects to decode
	for decoder.More() {
		var item []cartItem
		err := decoder.Decode(&item)
		if err != nil {
			panic(err)
		}
		cart = append(cart, item...)
	}

	return cart
}
