package main

import (
	"fmt"
	"io"
	"os"
)

func io_practice() {
	fmt.Println("Files")
	content := "Hello from Go!"
	file, err := os.Create("./fromString.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v cahracters\n ", length)
	defer file.Close() // very important to close the file

	defer readFile("./fromString.txt") // very important to defer this because the file needs to be closed before reading
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println("Data read from file: ", string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
