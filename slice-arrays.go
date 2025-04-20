package main

import "fmt"

func main() {
	var z []float64
	fmt.Printf("type of x %T\n",z) // this is slice
	var a [3]float64
	fmt.Printf("type of a %T\n", a) // this is array, as it's size is 3
	b := [...]string{"John", "Bob"} // this is also array, here compiler will count the no.of elements and defines it's size
	fmt.Println(b)
	fmt.Printf("type of string array %T\n", b)
	var c []int // this is slice
	fmt.Printf("len of slice %d capacity of slice %d\n", len(c) , cap(c))
	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:]
	fmt.Printf("len of slice %d capacity of slice %d\n", len(e) , cap(e))
	fmt.Println(e)
	e[1] = 'm'
	fmt.Println(e)
	fmt.Println(d)
	f := []int{1,3,5,7,9}
	g := f[:]
	// g[5] = 11 // this will raise an error index out of range [5] with length 5
	fmt.Println(g)
	fmt.Printf("len of f:%d capacity of f:%d\n", len(f),cap(f))
	f = append(f, 11,13)
	fmt.Printf("len of f:%d capacity of f:%d\n", len(f),cap(f))
	f = append(f, 13)
	fmt.Printf("len of f:%d capacity of f:%d\n", len(f),cap(f))
	slice1 := []int{1,2,3}
	var slice2 []int
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
	// here it will not copy into slice2 as slice2's capacity is 0
	slice3 := make([]int, 2)
	copy(slice3, slice1)
	fmt.Println(slice1, slice3)

}
