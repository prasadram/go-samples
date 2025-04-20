package main

import "fmt"

func main() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	test := (true && false) || (false && true) || !(false && false)
	fmt.Println("result:", test)
}
