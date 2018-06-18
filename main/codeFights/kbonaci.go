package main

import (
	"fmt"
	m "math/big"
)

//1 ≤ n ≤ 10^5
//1 ≤ k ≤ 10^3
func main() {
	fmt.Printf("\nkibonaciiii %s \n", kbonacci(10, 15))
}

func kbonacci(kParam int, n int) string {
	if n < kParam {
		return string(1)
	}
	window := make([]*m.Int, kParam+1)

	for i := 0; i < kParam; i++ {
		window[i] = m.NewInt(1)
	}
	window[kParam] = m.NewInt(int64(kParam))
	k := kParam
	//printListR(window)
	two := m.NewInt(2)
	/*
        for (int i = k; i <= n; i++) {
            int ix = i % k;
            BigInteger a = window[(i + k - 1) % k];
            BigInteger b = window[ix];
            BigInteger c = a.multiply(two).subtract(b);
            window[ix] = c;
        }

        return window[n % k].toString(10);
	*/

	for z := k; z <= n; z++ {
		ix := z % k
		a := window [(z+k-1)%k ]
		b := window [ix]
		window[ix].Mul(a, two)
		window[ix].Sub(window[ix], b)
		fmt.Println(window[ix].Text(10))
	}
	return window[n%(k-1)].Text(10)
}
func printListR(window []*m.Int) {
	for l := 0; l < len(window); l++ {
		fmt.Println(window[l].Text(10))
	}
	fmt.Println()
}
