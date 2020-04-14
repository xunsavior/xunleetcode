package bs

/*
Given a sorted array and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:
Input: [1,3,5,6], 5
Output: 2

Example 2:
Input: [1,3,5,6], 2
Output: 1

Example 3:
Input: [1,3,5,6], 7
Output: 4

Example 4:
Input: [1,3,5,6], 0
Output: 0
*/

// Time: O(logN)
// Space: O(1)
func searchInsert35(nums []int, target int) int {
	n := len(nums)
	// smaller than the min number
	if n == 0 || target < nums[0] {
		return 0
	}
	// greater than the max number
	if target > nums[n-1] {
		return n
	}

	start, end := 0, len(nums)-1

	for start < end {
		mid := start + (end-start)/2
		if nums[mid] < target {
			start = mid + 1
			continue
		}
		end = mid
	}

	return start
}
