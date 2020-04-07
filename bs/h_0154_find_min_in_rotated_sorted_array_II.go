package bs

/*
Company: Amazon(3); nil; Facebook(3), Uber(2), Adobe(2), Google(2)

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).

Find the minimum element.
The array may contain duplicates.

Example 1:
Input: [1,3,5]
Output: 1

Example 2:
Input: [2,2,2,0,1]
Output: 0

Note:
This is a follow up problem to Find Minimum in Rotated Sorted Array.
Would allow duplicates affect the run-time complexity? How and why?
*/

func findMin154(nums []int) int {
	start, end := 0, len(nums)-1
here:
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] > nums[end] {
			start = mid + 1
		} else if nums[mid] < nums[end] {
			end = mid
		} else {
			slice := nums[mid:end]
			for _, v := range slice {
				if v != nums[mid] {
					start = mid + 1
					continue here
				}
			}
			end = mid
		}
	}

	return nums[start]
}
