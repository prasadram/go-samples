package main

import "fmt"
import "time"

func main() {
c1 := make(chan string, 2)
go func() {
	c1 <- "Hello world"
	time.Sleep(time.Second * 2)
	c1 <- "Welcome To Buffered Channels"
	time.Sleep(time.Second * 2)
	c1 <- "is previous are read by consumer"
}()
	
	fmt.Println(<-c1)
	fmt.Println("length of channel ", len(c1))
	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println("Completed")
	
}
