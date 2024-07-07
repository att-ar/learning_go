package main

import (
	"fmt"
)


func methods() {
	poodle := Dog{"Poodle", 5, "Rawr!"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)

	poodle.Speak()
	poodle.Sound = "Woof!"
	poodle.Speak()

	poodle.SpeakThreeTimes()
	fmt.Printf("%+v\n", poodle)

	poodle.ChangeSound("Bark!")
	poodle.Speak()
}

// Dog is a struct
type Dog struct {
	Name string
	Age int
	Sound string
}

// Speak is a method of Dog
func (d Dog) Speak() { // (d Dog) is the receiver, that's how we define a method
	fmt.Println(d.Sound)
}

// SpeakThreeTimes is a method of Dog
func (d Dog) SpeakThreeTimes() {
	// d is a copy of the passed Dog object
	// so we can't change the original object
	// to do that we would need to use pointers instead
	d.Sound = fmt.Sprintf("%v! %v! %v!", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

// ChangeSound is a method of Dog
func (d *Dog) ChangeSound(sound string) {
	// d is a pointer to the passed Dog object
	// this overwites the original object
	d.Sound = sound
}