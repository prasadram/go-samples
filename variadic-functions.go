package main

import "fmt"

func main() {
	fmt.Println(sum(10,20))
	fmt.Println(sum(10,20,30))
}

func sum(args ...int) int {
	total := 0
	for _, value := range args {
		 total += value
	}
	return total
}
