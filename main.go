package main // Progams relatd to console based

import "fmt" // Library which contains functions related to format strings and print them in the console

func main() {
	// string
	var name string = "Isuru"
	var name2 string
	name3 := "Harischandra"
	fmt.Println(name, name2, name3)

	// int
	var num1 int = 234
	var num2 int8 = 21
	var num3 int16 = -3863
	var num4 int32 = 379322
	var num5 int64
	num6 := 3343
	fmt.Println(num1, num2, num3, num4, num5, num6)

	// uint (unsigned int - Doesn't contain negative values)
	var unum1 uint = 234
	var unum2 uint8 = 21
	var unum3 uint16 = 3863
	var unum4 uint32 = 379322
	var unum5 uint64
	fmt.Println(unum1, unum2, unum3, unum4, unum5)

	// float
	var fnum1 float32 = 29732.2324
	var fnum2 float64 = -3023902.23
	fnum3 := 323.23
	fmt.Println(fnum1, fnum2, fnum3)
}
