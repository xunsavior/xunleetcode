package main

import (
	"sort"
	"strconv"
	"strings"
)

func main() {
	largestNumber179([]int{3, 30, 34, 5, 9})
}
func largestNumber179(nums []int) string {
	strnums := make([]string, 0)
	for _, v := range nums {
		strnums = append(strnums, strconv.Itoa(v))
	}
	sort.Slice(strnums, func(i, j int) bool {
		return (strnums[i] + strnums[j]) > (strnums[j] + strnums[i])
	})
	numStr := strings.Join(strnums, "")
	numStr = strings.TrimLeft(numStr, "0")
	return numStr
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
