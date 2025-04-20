package main

import "fmt"

func main() {
	foo([]int{1,2,3,4}[0:1])	
}

func foo(bar []int) {
	fmt.Println(len(bar))
	fmt.Println(bar[0:3])
}
