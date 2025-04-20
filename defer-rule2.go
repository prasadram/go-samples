package main

import "fmt"

// Deferred function calls are executed in Last In First Out order after the surrounding function returns.

func main(){
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}
