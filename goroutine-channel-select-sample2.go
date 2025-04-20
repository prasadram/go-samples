package main

import "fmt"
import "time"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	
	go func() {
		for {
			select {
				case msg1 := <- c1:
				  fmt.Println(msg1)
				case msg2 := <- c2:
				  fmt.Println(msg2)
				case msg3 := <- time.After(time.Second):
				  fmt.Println("timeout", msg3)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
