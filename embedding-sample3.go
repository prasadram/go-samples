package main

import "fmt"

// Base
type Person struct {
	Name string
	Age int
}

func (p Person) Greet() {
	fmt.Println("Hello Person, my name is", p.Name)
}

// Embedding Person in Employee
type Employee struct {
	Person
	Company string
}

func (e Employee) Greet() {
	fmt.Println("Hello Employee, my name is", e.Name)
}

func main() {
	e := Employee{
		Person: Person{Name: "Bob", Age: 30},
		Company: "Meta",
	}
	e.Greet() // Hello Employee, my name is Bob 
	e.Person.Greet() // Hello Person, my name is Bob
}
