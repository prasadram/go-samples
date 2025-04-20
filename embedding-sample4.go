package main

import "fmt"

// Base
type Person struct {
	Name string
	Age int
}

func (p *Person) SetName(newName string) {
	p.Name = newName
}

// Embedding Person in Employee
type Employee struct {
	*Person
	Company string
}

func main() {
	p := &Person{Name: "Bob", Age: 30}
	e := Employee{
		Person: p,
		Company: "Meta",
	}
	// Accessing fields directly
	fmt.Println(e.Name)
	fmt.Println("Modify the name using e")
	e.SetName("John")
	fmt.Println("After modification name using e ", e.Name)
	fmt.Println("After modification name using p ", p.Name)
	fmt.Println("Modify the name using p")
	p.SetName("David")
	fmt.Println("After modification name using e ", e.Name)
	fmt.Println("After modification name using p ", p.Name)
}
