package main

//import "fmt"

func main() {
	var age int = 18
	if age > 18 {
	 print("you can vote\n")
	} else {
	 print("you can't vote\n")
	}
	// default is optional in switch
	switch age {
	  case 15: print("School\n")
	  case 18: print("College\n")
	}
}
