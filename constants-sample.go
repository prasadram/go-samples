package main

import "fmt"

const y string = "Welcome"
// z := "hello" // this will throw error as non-declaration statement outside function body
func main() {
	const x = "Hello World"
	fmt.Println(x)
	fmt.Println(y, "John")
	// fmt.Println(z)
}
