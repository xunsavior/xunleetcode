package bs

/*
Company: Amazon(4), Facebook(2); Microsoft(3), Google(2), Bloomberg(2)

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).

You are given a target value to search. If found in the array return true, otherwise return false.

Example 1:
Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true

Example 2:
Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false

Follow up:
This is a follow up problem to Search in Rotated Sorted Array, where nums may contain duplicates.
Would this affect the run-time complexity? How and why?
*/

func search81(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	left, right := 0, len(nums)-1
quiries:
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			// handle special case when the value of mid is equal to the right value
			for _, v := range nums[mid : right+1] { // [3, 1, 1, 1, 1], [2, 2, 2, 0, 2, 2]
				if v != nums[mid] {
					left = mid + 1
					continue quiries
				}
			}
			right = mid
		}
	}

	pivot := left
	left, right = 0, len(nums)-1
	if target >= nums[pivot] && target <= nums[right] {
		left = pivot
	} else {
		right = pivot - 1
	}

	for left <= right {
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}
