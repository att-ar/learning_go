package main

import (
	"fmt"
)

func pointers() {
	anInt := 42
	var p *int // * means it is a pointer
	fmt.Println("Pointers")
	fmt.Println(p)
	p = &anInt // & means the memory addres of anInt, not the value itself
	// so the value of p is the memory address of anInt

	fmt.Println(p)  // address of anInt
	fmt.Println(*p) // value of anInt
	fmt.Println(&p) // address of p

	value1 := 42.13
	pointer1 := &value1

	fmt.Println("Value 1:", *pointer1)

	*pointer1 = *pointer1 / 31 // dereferencing the pointer and changing the value (changes the value of value1 since we are modifying value stored in memory)

	fmt.Println("Pointer 1:", *pointer1)
	fmt.Println("Value 1:", value1)
}

// Referencing (&)
// Referencing is the act of obtaining the memory address of a variable.
// This is done using the & operator.
// When you place & before a variable,
// you're asking for the address in memory where that variable is stored.

// Dereferencing (*)
// Dereferencing is the act of accessing the value stored at a memory address pointed to by a pointer.
// This is done using the * operator placed before a pointer variable.

// Here's a quick summary:

// Declaration: var pointer1 *float64 declares pointer1 as a pointer to a float64.
// Dereferencing: *pointer1 accesses the float64 value that pointer1 points to.
// This dual use of the * symbol is a fundamental aspect of Go's syntax for working with pointers.
