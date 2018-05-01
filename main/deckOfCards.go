package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	start := time.Now()
	cards := make([]int, 52)
	for i := 1; i <= 52; i++ {
		cards[i-1] = i
	}
	var m = make(map[int32][]int)
	flag := 0
	for at := 0; at < 1000000000; at++ {
		if flag == 50000 {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			fmt.Printf(" %d  Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v \n", at, m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
			flag = 0
		}
		flag++
		cards := cards[:]

		shuffleArray(cards)
		hash := getHash(cards)
		fmt.Printf("%d %v \n", hash, cards)
		collision, ok := m[int32(hash)]
		if ok {
			fmt.Printf("%d \n %v  \n %v \n", hash, collision, cards)
			break
		}
		m[int32(hash)] = cards
	}
	//fmt.Print(cards)
	elapsed := time.Since(start)
	fmt.Println(elapsed)

}
func getHash(cards []int) int {
	hash := int(0)
	for j := 0; j < 10; j++ {
		hash += cards[j] * (j * 10)
	}
	return hash
}

func shuffleArray(ar []int) {
	for i := len(ar) - 1; i > 0; i-- {
		index := rand.Intn(i + 1)
		// Simple swap
		ar[i], ar[index] = ar[index], ar[i]
	}
}
