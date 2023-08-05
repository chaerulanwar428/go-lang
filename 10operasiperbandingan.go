package main

import "fmt"

func main() {
	var name1 = "Chaerul"
	var name2 = "Anwar"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 100

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}