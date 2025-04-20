package main

import "fmt"

// Base
type Person struct {
	Name string
	Age int
}

func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

// Embedding Person in Employee
type Employee struct {
	Person
	Company string
}

func main() {
	e := Employee{
		Person: Person{Name: "Bob", Age: 30},
		Company: "Meta",
	}
	e.Greet()
	// Even though Greet() is defined in Person, we can call e.Greet() on Employee because Person is embedded.
}
