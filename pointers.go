package main

import "fmt"


func main() {
	i := 7
	// Send the pointer using ampersand
	inc(&i)
	fmt.Println(i)
}

// Function that accept pointers
func inc(x *int) {
	// Dereference the pointer
	*x++
}