package leetcodetest

import (
	"leetcode/collections"
	"testing"
)

type ThreeSum struct {
	nums      []int
	expResult [][]int
}

var threeSumTestCases = []*ThreeSum{
	&ThreeSum{
		nums: []int{-1, 0, 1, 2, -1, -4},
		expResult: [][]int{
			[]int{-1, 0, 1},
			[]int{-1, -1, 2},
		},
	},
	&ThreeSum{
		nums: []int{0, 0, 0},
		expResult: [][]int{
			[]int{0, 0, 0},
		},
	},
}

func TestThreeSum(t *testing.T) {
	for _, v := range threeSumTestCases {
		equals(t, v.expResult, collections.ThreeSum(v.nums))
	}
}
