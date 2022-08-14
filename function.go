package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var x int = 3
	var y int = 7

	var total int = sum(x, y)

	fmt.Println(total)

	var z float64 = -0.1

	z_sqrt, err := sqrt(z)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(z_sqrt)
	}
}

// Function that returns a value
func sum(x int, y int) int {
	return x + y
}

// Function that also handles the error
func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("NegativeRootError, cannot calculate the root of negative integer")
	}

	return math.Sqrt(num), nil
}
