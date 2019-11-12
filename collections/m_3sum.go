package collections

import "math"

/*
	Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
	Find all unique triplets in the array which gives the sum of zero.

	The solution set must not contain duplicate triplets.

	Example:
	Given array nums = [-1, 0, 1, 2, -1, -4],
	A solution set is:
	[
	[-1, 0, 1],
	[-1, -1, 2]
	]
*/

// ThreeSum ...
func ThreeSum(nums []int) [][]int {
	out := [][]int{}
	negNums, posNums := make([]int, 0, len(nums)), make([]int, 0, len(nums))
	var (
		negIndex int
		posIndex int
		negMap   = map[int]int{} //key: value of element; value: the number of elements with the same value
		posMap   = map[int]int{}
	)

	// divide original numbers into positive (include 0) and negative slices
	// store both slices into hash map
	for _, v := range nums {
		if v < 0 {
			negMap[v] = negMap[v] + 1
			negNums = append(negNums, v)
			negIndex++
		} else {
			posMap[v] = posMap[v] + 1
			posNums = append(posNums, v)
			posIndex++
		}
	}

	// special case: if the positive map has more than 3 zeros, then we add [0, 0, 0]
	if numZeroIndexes, ok := posMap[0]; ok && numZeroIndexes >= 3 {
		out = append(out, []int{0, 0, 0})
	}

	// create a map to check if the negative/positive pair exists
	var existedPair = make(map[[2]int]bool)
	// pick a negative number from negMap, and search for the other two positive numbers from  pos slice
	for k := range negMap {
		sum := -k
		for _, num := range posNums {
			c := sum - num
			numIndexes, ok := posMap[c]
			if ok && numIndexes >= 1 {
				max := math.Max(float64(c), float64(num))
				min := math.Min(float64(c), float64(num))
				arrayK := [2]int{int(min), int(max)}
				if existedPair[arrayK] == false {
					existedPair[arrayK] = true
					if num == c && numIndexes == 1 {
						continue
					}
					out = append(out, []int{k, int(min), int(max)})
				}
			}
		}
	}
	// pick a positive number from negMap, and search for the other two negative numbers from neg slice
	for k := range posMap {
		sum := -k
		for _, num := range negNums {
			c := sum - num
			numIndexes, ok := negMap[c]
			if ok && numIndexes >= 1 {
				max := math.Max(float64(c), float64(num))
				min := math.Min(float64(c), float64(num))
				arrayK := [2]int{int(min), int(max)}
				if existedPair[arrayK] == false {
					existedPair[arrayK] = true
					if num == c && numIndexes == 1 {
						continue
					}
					out = append(out, []int{int(min), int(max), k})
				}
			}
		}
	}

	return out
}
