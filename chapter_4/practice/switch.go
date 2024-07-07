package main

import (
	"fmt"
	"math/rand"
	"time"
)

func switch_fn() {
	rand.NewSource(time.Now().Unix())
	// fmt.Println("Day", dow) // no longer need this since I am declaring dow in the switch statement

	var result string
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "Sunday"
	case 2:
		result = "Monday"
	case 3:
		result = "Tuesday"
	case 4:
		result = "Wednesday"
	case 5:
		result = "Thursday"
	case 6:
		result = "Friday"
		// fallthrough // this prevents the switch statement from breaking, so it will continue to the next case (for no reason in this situation)
	case 7:
		result = "Saturday"
	default:
		result = "Invalid day"
	}
	fmt.Println(result)

}