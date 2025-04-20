package main

import "fmt"


// Base
type Person struct {
	Name string
	Age int
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
	// Accessing fields directly
	fmt.Println(e.Name)
	fmt.Println(e.Age)
	fmt.Println(e.Company)
}
