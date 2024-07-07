package main

import (
	"fmt"
)

type cartItem struct{
    name string
    price float64
    quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []cartItem) float64 {
    var result float64
	for _, item := range cart {
		// every item is of type cartItem
		result += item.price * float64(item.quantity)
	}
	return result
}

func main() {
	cart := []cartItem{
		{"Shoes", 42.50, 2},
		{"Socks", 5.00, 6},
		{"Pants", 28.75, 1},
	}
	fmt.Println("Total:", calculateTotal(cart))
}