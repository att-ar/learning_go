package main // to serve as starting point of the program

import (
	"fmt"
	"time"
)

// declaring a constant available to all functions in the package
const aConst int64 = 909

// lower case name for private functions/methods/variables etc
// upper case name for public functions/methods/variables etc

func main() {
	n := time.Now()
	fmt.Println("I recorded this video at", n)

	t := time.Date(2020, time.November, 10, 20, 10, 0, 0, time.UTC)
	fmt.Println("Go launched at ", t)
	fmt.Println((t.Format(time.ANSIC)))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 20:10:00 2020")
	fmt.Printf("The Type of parsedTime is %T\n", parsedTime)
}

// func main() {
// 	i1, i2, i3 := 12, 45, 68
// 	intSum := i1 + i2 + i3
// 	fmt.Println("Integer sum:", intSum)

// 	f1, f2, f3 := 23.5, 65.1, 76.3
// 	floatSum := f1 + f2 + f3
// 	fmt.Println("Float sum:", floatSum)

// 	// equals not colon equals because it is already initialized
// 	floatSum = math.Round(floatSum * 100) / 100
// 	fmt.Println("Rounded float sum:", floatSum)

// 	circleRadius := 15.5
// 	circumference := circleRadius * 2 * math.Pi
// 	fmt.Printf("Circumference: %.2f\n", circumference)
// }

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter text: ")
// 	input, _ := reader.ReadString('\n') // characters are wrapped in single quotes, strings are wrapped in double quotes
// 	fmt.Println("You entered:", input)

// 	fmt.Print("Enter a number: ")
// 	numInput, _ := reader.ReadString('\n')
// 	aFloat, err := strconv.ParseFloat((strings.TrimSpace(numInput)), 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Value of number:", aFloat)
// 	}
// }

// func main() {
// 	// explicit typing
// 	var aString string = "This is Go!"
// 	fmt.Println(aString)
// 	fmt.Printf("The variable's type is %T\n", aString)

// 	var anInteger int = 42
// 	fmt.Println(anInteger)

// 	var defaultInt int

// 	fmt.Println(defaultInt)

// 	// implicit typing
// 	var anotherString = "This is also Go!"
// 	fmt.Println(anotherString)
// 	fmt.Printf("The variable's type is %T\n", anotherString)

// 	var anotherInt = 5
// 	fmt.Println(anotherInt)
// 	fmt.Printf("The variable's type is %T\n", anotherInt)

// 	myString := "This is a colonated variable string, only works inside a function"
// 	fmt.Println(myString)
// 	fmt.Printf("The variable's type is %T\n", myString)

// 	fmt.Println("A constant", aConst)
// }