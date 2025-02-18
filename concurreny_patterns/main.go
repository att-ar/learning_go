package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println(("Log: working..."))
		}
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		default:
			fmt.Println("Waiting for messages")
			time.Sleep(500 * time.Millisecond) // Added to prevent a tight loop
		}
	}

	done := make(chan bool)

	go doWork(done)

	time.Sleep(1 * time.Second)

	close(done)

}
