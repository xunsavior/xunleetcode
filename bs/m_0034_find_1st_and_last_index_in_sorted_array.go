package bs

/*
Company: Facebook(24), Google(10), Uber(9), LinkedIn(6), Apple(5)

Given an array of integers nums sorted in ascending order, find the starting
and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Example 2:
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
*/

func searchRange34(nums []int, target int) []int {
	res := []int{-1, -1}
	left, right, targetIndex := 0, len(nums)-1, -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			targetIndex = mid
			break
		}
	}

	if targetIndex == -1 {
		return res
	}

	res[0], res[1] = targetIndex, targetIndex

	left, right = targetIndex, targetIndex

	for left >= 0 {
		if nums[left] == target {
			res[0] = left
			left--
			continue
		}
		break
	}

	for right <= len(nums)-1 {
		if nums[right] == target {
			res[1] = right
			right++
			continue
		}
		break
	}

	return res
}
