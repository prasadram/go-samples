package main

import "fmt"

// A deferred function’s arguments are evaluated when the defer statement is evaluated(not during execution).

func main() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
