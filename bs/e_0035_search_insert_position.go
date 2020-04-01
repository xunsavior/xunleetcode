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

func searchInsert35(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	start, end, mid := 0, len(nums)-1, (len(nums)-1)/2

	if target < nums[start] {
		return 0
	}

	if target > nums[end] {
		return len(nums)
	}

	res := 0
	for start <= end {
		midVal := nums[mid]
		if target > midVal {
			start = mid + 1
			if midVal < nums[start] {
				res = mid
				break
			}
		} else if target < midVal {
			end = mid - 1
			if midVal > nums[end] {
				res = mid
				break
			}
		} else {
			res = mid
			break
		}
		mid = (start + end) / 2
	}

	return res
}
