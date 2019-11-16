package main

import (
	"fmt"
	"log"
	"time"
)

//nums "leetcode/nums"

func main() {
	sampleDPDemo()
}

/*
	f(0) = 0
	f(1) = 1
	f(2) = f(1) + f(0)
	...
	f(n) = f(n-1) + f(n-2)
*/

func sampleDPDemo() {
	n := 30
	cache := make([]int, n+1)
	start := time.Now()
	resf := f(n, cache)
	duration := time.Since(start)
	fmt.Printf("Result is %d and f() took %d\n", resf, duration.Nanoseconds())

	start = time.Now()
	resf = fSlow(n)
	duration = time.Since(start)
	log.Printf("Result is %d and fSlow() took %d", resf, duration.Nanoseconds())

}

// Fibonacci
func f(n int, cache []int) int {
	if cache[n] != 0 {
		return cache[n]
	}
	if n < 3 {
		return n
	}
	res := f(n-2, cache) + f(n-1, cache)
	cache[n] = res
	return res
}

func fSlow(n int) int {
	if n < 3 {
		return n
	}
	return fSlow(n-2) + fSlow(n-1)
}
