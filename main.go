package main

import (
	nums "leetcode/nums"
	"log"
)

func main() {
	n1 := []int{1, 2, 3}
	n2 := n1[0:3]
	log.Println(n2)
	nums.TestMaxArea()
}
