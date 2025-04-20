package main

import "fmt"

func main() {
	xPtr := new(int) // new takes a type as an argument, allocates enough memory to fit a value of that type and returns a pointer to it.
	c(xPtr)
	fmt.Println(*xPtr)
}

func c(xPtr *int) {
	*xPtr = 1
}
