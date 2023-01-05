package main

import "fmt"

type Person struct {
	Name string
}

type Android struct {
	Person
	Name string
}

func main() {

	android := new(Android)
	android.Person = Person{Name: "Rakka"}

	fmt.Println(android.Person.Name)
}
