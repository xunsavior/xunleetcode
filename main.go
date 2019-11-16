package main

import (
	"log"
	"reflect"
)

//nums "leetcode/nums"

func main() {
	a := []int{1, 3, 3}
	b := []int{1, 2, 3}
	log.Println(reflect.DeepEqual(a, b))
}

func addItem(nums []int, newValue int) []int {
	nums = append(nums, newValue)
	return nums
}
