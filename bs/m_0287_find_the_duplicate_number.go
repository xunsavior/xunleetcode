package bs

/*
Company: Amazon(7), Microsoft(4), Google(4), Facebook(3), Apple(2)

Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive),
prove that at least one duplicate number must exist. Assume that there is only one duplicate
number, find the duplicate one.

Example 1:

Input: [1,3,4,2,2]
Output: 2
Example 2:

Input: [3,1,3,4,2]
Output: 3
Note:

You must not modify the array (assume the array is read only).
You must use only constant, O(1) extra space.
Your runtime complexity should be less than O(n2).
There is only one duplicate number in the array, but it could be repeated more than once.
*/

// value boundary
func findDuplicate287(nums []int) int {
	left, right := 1, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		count1, count2 := 0, 0
		for _, v := range nums {
			if v < mid {
				count1++
			} else if v == mid {
				count2++
			}
		}

		if count2 > 1 {
			return mid
		}

		if count1 < mid-1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
