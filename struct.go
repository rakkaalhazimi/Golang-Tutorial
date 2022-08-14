package main

import "fmt"

// Struct is collections of field
type Person struct {
	name string
	age int
}

// Method of the struct
func (p *Person) Greet() {
	fmt.Printf("Hello my name is %s. I'm %d years old", p.name, p.age)
}

func main() {
	// Initialize struct instance with predefined fields
	p1 := Person{name: "Rakka", age: 20}

	// Or leave it blank
	p2 := Person{}

	fmt.Println("P1 name: ", p1.name)
	fmt.Println("P2 name: ", p2.name)

	// Call method
	p1.Greet()
}
