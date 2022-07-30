package main

import "fmt"

func main() {
	/* Arrays */
	var numbers1 [3]int8 = [3]int8{10, 20, 30}
	fmt.Println(numbers1, len(numbers1))

	var numbers2 = [2]int16{40, 50}
	fmt.Println(numbers2, len(numbers2))

	numbers3 := [5]int32{60, 70, 80, 90, 100}
	fmt.Println(numbers3, len(numbers3))

	numbers4 := [5]int64{2, 4, 6, 8}
	fmt.Println(numbers4, len(numbers4))

	names := [3]string{"Isuru", "Vihan", "Harischandra"}
	names[1] = "VIHAN"
	fmt.Println(names, len(names))

	/* Slices */
	var scores = []int{100, 200, 300}
	scores = append(scores, 400, 500) // Appending
	fmt.Println(scores, len(scores))

	// Slice ranges
	rangeOne := scores[:3]
	rangeTwo := scores[1:4]
	rangeThree := scores[2:]
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)
}
