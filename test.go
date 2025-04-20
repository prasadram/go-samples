package main

import "fmt"

func main() {
	var x string = "hello"
	var y string = "hello"
	fmt.Println("x address: ", &x, "y address: ", &y)
	fmt.Println(x == y)
	var z []float64
	fmt.Printf("type of x %T\n",z)
	var a [3]float64
	fmt.Printf("type of a %T\n", a)
	b := [...]string{"John", "Bob"}
	fmt.Println(b)
	fmt.Printf("type of string array %T\n", b)
	var c []int
	fmt.Printf("len of slice %d capacity of slice %d\n", len(c) , cap(c))
	d := [6]string{"a","b","c","d","e","f"}
	fmt.Println(d[2:5])
	e := 1.5
  	square(&e)
	fmt.Println(e)
}

func square(x *float64) {
  *x = *x * *x
}

