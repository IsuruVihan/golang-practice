package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "Hello! I am Isuru!"
	fmt.Println(strings.Contains(greeting, "Hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "ello", "Hi")) // This doesn't change the original string
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "am"))
	fmt.Println(strings.Split(greeting, " "))

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages) // This changes the original order of integers
	fmt.Println(ages)

	index := sort.SearchInts(ages, 45)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names) // This changes the original order of string
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "dvdsvssdv"))
}
