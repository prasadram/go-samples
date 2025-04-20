package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		fmt.Println(i);
	}
	// while type 
	fmt.Println("implementing while");
	i := 1
	for i < 5 {
		fmt.Println(i);
		i++;
	}
	fmt.Println("nested loops");
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			fmt.Println("*");
		}
			fmt.Println();
	}

}
