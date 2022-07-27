package main // Progams relatd to console based

import "fmt" // Library which contains functions related to format strings and print them in the console

func main() {
	// Print() -> Not adding a new line automatically
	fmt.Print("Hello \n")
	fmt.Print("World!")

	// Println() -> Adding a new line automatically
	fmt.Println("Hello")
	fmt.Println("World!")

	// Printf() -> Formatted printing
	name := "Isuru"
	age := 23
	points := 232.23
	fmt.Printf("My name is %v and my age is %v\n", name, age) // Variables
	fmt.Printf("My name is %q and my age is %q\n", name, age) // Quotes (Only works for strings)'
	fmt.Printf("Type of 'age' variable: %T\n", age)           // Type
	fmt.Printf("Points scored: %f\n", points)                 // Floats
	fmt.Printf("Points scored: %0.2f\n", points)              // Floats - Limiting decimal points

	// Sprint() -> Save print
	var str = fmt.Sprintf("My name is %v and my age is %v\n", name, age)
	fmt.Print("Sentence: ", str)
}
