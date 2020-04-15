package bs

/*
Company: Facebook(30), Amazon(22), Microsoft(10), Bloomberg(9), Oracle(8)

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Example 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
*/

// time: O(logN)
// space: O(1)
func search33(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	// find the pivot
	for start < end {
		mid := start + (end-start)/2
		midVal := nums[mid]
		if midVal > nums[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}

	pivot := start
	start, end = 0, len(nums)-1
	if target >= nums[pivot] && target <= nums[end] {
		start = pivot
	} else {
		end = pivot - 1
	}

	for start <= end {
		mid := start + (end-start)/2
		midVal := nums[mid]
		if target < midVal {
			end = mid - 1
		} else if target > midVal {
			start = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
