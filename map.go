package main

import "fmt"

func main() {
	city_code := make(map[string]int)

	// Create key - value association
	city_code["jakarta"] = 1
	city_code["semarang"] = 2
	city_code["surabaya"] = 3

	// Delete key
	delete(city_code, "semarang")

	// Change the value of the key
	city_code["jakarta"] = 4

	fmt.Println(city_code)
}
