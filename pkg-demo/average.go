package main

import "fmt"
import "example.com/mymath"
//import "../go-samples/mymath"

func main() {
	xs := []float64{1,2,3,4}
	avg := mymath.Average(xs)
	fmt.Println(avg)	
}
