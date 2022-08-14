package main

import (
	"basic/greet"
	"fmt"
)

func main() {
	names := []string{"combro", "tenyom", "diwut"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	greet.Greet("rakka")
}
