package main

import (
	"fmt"
	"sort"
)

func main() {
	test([]int{3, 2, 1})
}

func test(nums []int) {
	sort.Ints(nums)
	fmt.Println(nums)
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
