package main

import "fmt"

func main() {
	defer first()
	second()
}
func first() {
	fmt.Println("1st")
}
func second(){
	fmt.Println("2nd")
}

