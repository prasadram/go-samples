package main

import "fmt"

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recover from ", r)
	}
}
func fullName(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("runtime error: firstname cannot be null")
	}
	if lastName == nil {
		panic("runtime error: lastname cannot be null")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("defer call in main")
	firstName := "John"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}
