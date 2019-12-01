package main

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

}

func addMap(nums []int, s *[][]int) {
	nums = append(nums, 3)
	*s = append(*s, nums)
	nums = append(nums, 4)
	nums = append(nums, 5)
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
