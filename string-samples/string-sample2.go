package main

import "fmt"

func main() {
	fmt.Println("Hello\tWorld")
	fmt.Println("Hello\nWorld")
	fmt.Println(`Hello\nWorld`)
	fmt.Println(`Hello
		World`)
}
