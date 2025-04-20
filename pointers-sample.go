package main

import "fmt"

func main() {
	x := 5
	c(&x)
	fmt.Println(x)
}

func c(xPtr *int) {
	*xPtr = 0
}
