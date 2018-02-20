package main

import (
	f "fmt"
	"time"
)

func main() {
	f.Println("starting the all   " + time.Now().String())
	comm := make (chan string)
	go go1("thread1",12, comm)
	go go1("thread2",12, comm)
	go printIt(comm)

	var input string
	f.Scanln(&input)
}

func printIt(c <-chan string) {
	for {
		f.Println(<-c)
	//	time.Sleep(time.Second)
	}
}

func go1(name string,a int, comm chan<-  string)  {
	f.Println(name + " start")
	for i:=1;i<a;i=i+2   {
		comm <- name + " " + string(i)
	}
}