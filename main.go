package main

import "log"

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
	m := map[string]int{}
	//addMap(m)
	log.Println(m)
}

func addMap(m map[string]int) {
	m["1"] = 1
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
