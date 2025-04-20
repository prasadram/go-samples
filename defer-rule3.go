package main

import "fmt"

// Deferred functions may read and assign to the returning function’s named return values.

func main(){
	fmt.Println(c())
}

// a deferred function increments the return value i after the surrounding function returns. Thus, this function returns 2:
func c() (i int) {
    defer func() { i++ }()
    return 1
}

