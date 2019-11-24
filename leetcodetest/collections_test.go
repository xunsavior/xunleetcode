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

type ThreeSumClosest struct {
	nums      []int
	target    int
	expResult int
}

var threeSumClosestTestCases = []*ThreeSumClosest{
	&ThreeSumClosest{
		nums:      []int{-1, 2, 1, -4},
		target:    1,
		expResult: 2,
	},
	&ThreeSumClosest{
		nums:      []int{1, 1, 1, 0},
		target:    -100,
		expResult: 2,
	},
	&ThreeSumClosest{
		nums:      []int{0, 2, 1, -3},
		target:    1,
		expResult: 0,
	},
}

// func TestThreeSumClosest(t *testing.T) {
// 	for _, v := range threeSumClosestTestCases {
// 		equals(t, v.expResult, collections.ThreeSumClosest(v.nums, v.target))
// 	}
// }
