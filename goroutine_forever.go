package main

import (
	"fmt"
	"time"
)

func work(id int) {

	for {
		fmt.Printf("Worker #%d working...\n", id)
		time.Sleep(time.Second * time.Duration(2))
	}
}

func main() {
	go work(1)
	go work(2)

	var input string
	fmt.Scanln(&input)
}
