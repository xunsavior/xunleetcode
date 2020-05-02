package array

import "sort"

/*
Company: Amazon(2), Uber(2); Google(3); Refin(2), Bloomberg(2)

Given an integer array, you need to find one continuous subarray that
if you only sort this subarray in ascending order, then the whole
array will be sorted in ascending order, too.

You need to find the shortest such subarray and output its length.

Example 1:
Input: [2, 6, 4, 8, 10, 9, 15]
Output: 5
Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
Note:
Then length of the input array is in range [1, 10,000].
The input array may contain duplicates, so ascending order here means <=.
*/

func findUnsortedSubarray581(nums []int) int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)

	start, end := -1, -1

	for i := range nums {
		if nums[i] != sortedNums[i] {
			if start == -1 {
				start = i
			} 
			end = i
		} 
	}

	return end - start
}
