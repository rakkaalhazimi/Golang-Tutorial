package main

import "fmt"

func main() {
	// Create an array with "var - varname - [number]" combination
	var x [5]int

	// Use := and {} if you want to define the content inside it
	y := [5]int{1, 2, 3, 4, 5}

	// Array of fixed size is inconvenient, so we can create slice instead
	// doing so, we gained access to append method
	z := []int{1, 2, 3}
	z = append(z, 2)

	fmt.Println(x, y, z)
}
