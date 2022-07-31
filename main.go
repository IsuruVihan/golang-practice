package main

import "fmt"

func main() {
	// While loop
	x := 0
	for x < 5 {
		fmt.Println("Value of x is: ", x)
		x++
	}

	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is: ", i)
	}

	// For-each loops
	names := []string{"Isuru", "Vihan", "Harischandra", "Royal", "Colombo"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names { // When we need both index, and value
		fmt.Println("The value at index", index, "is", value)
	}

	for _, value := range names { // When we only need value
		fmt.Println("The value is", value)
	}
}
