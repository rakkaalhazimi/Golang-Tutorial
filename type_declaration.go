package main

import "fmt"

func main() {

	type NoKTP string
	type Married bool

	var MyID NoKTP = "12345678"
	var MarriedStatus Married = false
	fmt.Println(MyID)
	fmt.Println(MarriedStatus)

}
