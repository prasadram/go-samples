package main

import "fmt"

func main() {
	var evenNum[5] int
	evenNum[0] = 2
	evenNum[1] = 4
	evenNum[2] = 6
	evenNum[3] = 8
	evenNum[4] = 10
	print(evenNum[1])
	print("\n")
	fmt.Println(evenNum[2])
	//evenNum = append(evenNum, 12) // this will not work as append will wok on slies only
	oddNum := [3]int{1,3,5}
	oddNum[0] = 1
	fmt.Println(oddNum[2])
	fmt.Println("loops for arrays")
	for _, value := range evenNum {
	  fmt.Println(value)
	}
	fmt.Println("loops for arrays using indes/iterator")
	for i, value := range oddNum {
	  fmt.Println(value, i)
	}
	fmt.Println("slicing")
	numSlice := []int{5,4,3,2,1}
	sliced := numSlice[3:5]
	fmt.Println(sliced)
	backSlice := numSlice[:5]
	fmt.Println(backSlice)
	
	anotherSlice := make([]int, 5, 10)
	copy(anotherSlice,  numSlice)
	fmt.Println(anotherSlice)
}
