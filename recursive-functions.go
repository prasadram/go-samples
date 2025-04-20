package main

import "fmt"

func main() {
	fmt.Println("Factorial : ", factorial(5))
}

func factorial(a int) int {
        if a == 0 {
		return 1
	}

	return a * factorial(a-1)
}
