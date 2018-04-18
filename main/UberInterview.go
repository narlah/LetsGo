package main

//Our server is calling the external API using call_api that is doing the HTTP call.
//The API provider charges us if we make more than 10 calls per second.
// How would you implement something that keeps us below the limit?

import (
	"time"
	"fmt"
	"strconv"
)

const Limit = 10

type transaction struct {
	data string
	time time.Time
}

func main() {
	var lastTransactions []transaction
	go cleanQueue(&lastTransactions)
	for i := 0; i <= 100; i++ {
		receive("Transaction with id : "+strconv.Itoa(i), &lastTransactions)
		if i < 50 {
			time.Sleep(101 * time.Millisecond)
		}
	}

}
func receive(transMsg string, lastTransactions *[]transaction) {
	currentTime := time.Now()
	lenArr := len(*lastTransactions)
	if lenArr > Limit {
		tenBack := (*lastTransactions)[lenArr-Limit]
		fmt.Printf("diff with the last 10 %v  \n", getMilseconds(currentTime)-getMilseconds(tenBack.time))

		if getMilseconds(currentTime)-getMilseconds(tenBack.time) < 1000 {
			fmt.Printf("WARNING you are going too fast")
		}
	}
	fmt.Printf("Send msg %s \n", transMsg) //sending it
	*lastTransactions = append(*lastTransactions, transaction{transMsg, currentTime})
}
func cleanQueue(lastTransactions *[]transaction) {
	for i := 0; i < 10; i++ {
		fmt.Print(strconv.Itoa(len(*lastTransactions)) + " =>")
		if len(*lastTransactions) > 26 {
			*lastTransactions = (*lastTransactions)[14:]
		}
		fmt.Println(len(*lastTransactions))
		time.Sleep(1000 * time.Millisecond)
	}
}
func (t transaction) String() string {
	return t.data + "  current time : " + strconv.FormatInt(getMilseconds(t.time), 10)
}

func getMilseconds(cTime time.Time) int64 {
	return cTime.UnixNano() / int64(time.Millisecond)
}
