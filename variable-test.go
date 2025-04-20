package main

import "fmt"

func main(){
	var a int = 5;
	var b float32 = 4.32;
	const pi float64 = 3.141
	fmt.Println(pi);
	fmt.Println(a);
	fmt.Println(b);
	
	x, y := 34, 35
	fmt.Println(x);
	fmt.Println(y);
	fmt.Println(b);
	fmt.Println(x, ",", y);
	
	isbool := true;
	hate := false;
	fmt.Println(isbool && hate);
	fmt.Println(isbool || hate);
	fmt.Println(!isbool );

}
