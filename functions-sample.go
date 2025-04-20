package main

import "fmt"

func main() {
	a, b := testReturn()
	fmt.Println(a)
	fmt.Println(b)
}

func testReturn() (x int, z bool) {
	fmt.Println("test return called")
	x = 2
	return x, true
}
