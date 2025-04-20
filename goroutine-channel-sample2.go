package main

import(
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i:=0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		//msg := <- c
		//fmt.Println(msg + " " + time.Now().Format("15:04:05"))
		// above can be achieved like this as well
		fmt.Println(<- c + " " + time.Now().Format("15:04:05"))
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
