package main

import (
	"fmt"
	"time"
)

/*
Да се създадат три поцеса, базирани на C/C++ програми, комуникиращи през обща памет с размер едно цяло число.
Първият процес генерира числата от 1 до 100 и ги записва в обща памет.
Останалите процеси четат от общата памет, като формират съответно сумите на четните и нечетните числа и ги визуализират.
*/

func main() {
	ch := make(chan int, 1)
	go generate(ch)
	go consumeEven(ch)
	go consumeOdd(ch)
	time.Sleep(time.Second * 10000)
}

func generate(ch chan<- int) {
	for i := 1; i <= 102; i++ {
		ch <- i
	}
}

func consumeOdd(ch chan int) {
	sum := 0
	for {
		in := <-ch
		if in > 100 {
			break
		}
		if in%2 == 0 {
			ch <- in
			continue
		}
		sum += in
	}
	fmt.Printf("Odds sum : %d", sum)
}

func consumeEven(ch chan int) {
	sum := 0
	for {
		in := <-ch
		if in > 100 {
			break
		}
		if in%2 != 0 {
			ch <- in
			continue
		}
		sum += in
	}
	fmt.Printf("Odds sum : %d", sum)
}
