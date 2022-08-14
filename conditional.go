package main

import "fmt"

func main() {
	var x int
	x = 7

	if x < 6 {
		fmt.Println("Less than 6")
	} else if x > 6 {
		fmt.Println("More than 6")
	} else {
		fmt.Println("Equal 6")
	}

}
