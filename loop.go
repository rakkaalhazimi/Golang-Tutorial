package main

import "fmt"

func main() {
	// Regular while loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Regular while loop
	var j int = 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Loop through array
	array := []int{1, 2, 3, 4, 5}
	for index, value := range array {
		fmt.Println(index, value)
	}

	// Loop through map
	fruit_map := make(map[string]int)
	fruit_map["apple"] = 1
	fruit_map["blueberry"] = 2

	for key, value := range fruit_map {
		fmt.Println(key, value)
	}

}
