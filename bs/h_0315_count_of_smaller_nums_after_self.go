package bs

import "sort"

/*
Company: Google(11), Amazon(6), Microsoft(3); Apple(2), Bloomberg(2)

You are given an integer array nums and you have to return a new counts array.
The counts array has the property where counts[i] is the number of smaller
elements to the right of nums[i].

Example:

Input: [5,2,6,1]
Output: [2,1,1,0]
Explanation:
To the right of 5 there are 2 smaller elements (2 and 1).
To the right of 2 there is only 1 smaller element (1).
To the right of 6 there is 1 smaller element (1).
To the right of 1 there is 0 smaller element.
*/

// Fenwick tree
func countSmaller315(nums []int) []int {
	// reverse the nums
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	uniqueNums, uniqueMap := []int{}, make(map[int]bool)
	// make a sorted unique slice
	for _, v := range nums {
		if !uniqueMap[v] {
			uniqueMap[v] = true
			uniqueNums = append(uniqueNums, v)
		}
	}
	sort.Ints(uniqueNums)
	// create a function get the previous sum of index
	prevSumFunc := func(input []int, index int) int {
		sum := 0
		for i := range input {
			if i == index {
				break
			}
			sum += input[i]
		}
		return sum
	}

	// create a ranked slice for reversed nums
	rankedNums, freq, res := make([]int, 0, len(nums)), make([]int, 0, len(uniqueNums)+1), make([]int, 0, len(nums))
	for i, v := range nums {
		rank := helper315(uniqueNums, v)
		rankedNums = append(rankedNums, rank)
		freq[rank]++
		res[len(nums)-1-i] = prevSumFunc(freq, rank)
	}

	return res
}

// return the ranking
func helper315(nums []int, target int) int {
	start, end := 0, len(nums)-1

	for start <= end {
		mid := start + (end-start)/2
		if target < nums[mid] {
			end = mid - 1
		} else if target > nums[mid] {
			start = mid + 1
		} else {
			return mid + 1
		}
	}

	return -1
}
