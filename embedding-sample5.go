package main

import "fmt"

type Speaker interface {
	Speak()
}

type Greeter interface {
	Greet()
}

// New interface that embeds both Speaker and Greeter
type Communicator interface {
	Speaker
	Greeter
}

type Person struct {
	Name string
	Age int
}

func (p Person) Speak() {
	fmt.Println("Speaking...")
}

func (p Person) Greet() {
	fmt.Println("Hello i am ", p.Name)
}


func main() {
	var c Communicator = Person{Name: "John", Age: 30}
	c.Greet()
	c.Speak()
}
