package main

import "fmt"

func main() {
	pos, neg := add(), add()
	for i:=0; i <= 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

func add() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum 
	}
}
