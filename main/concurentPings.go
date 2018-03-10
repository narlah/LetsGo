package main

import (
	"fmt"
	"time"
)

func pinger(c chan string, pingerVersion string) {
	for i := 0; ; i++ {
		c <- "ping from " + pingerVersion

	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c, "first")
	go pinger(c, "second")
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
