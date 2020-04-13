package bs

import "sort"

/*
Company: Facebook(22), Oracle(5), Amazon(3), Microsoft(2), LinkedIn(2)
Given two arrays, write a function to compute their intersection.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]

Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]

Note:
Each element in the result must be unique.
The result can be in any order.
*/

func intersection349(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 || n2 == 0 {
		return nil
	}

	var unsortedArray, sortedArray []int
	if n1 < n2 {
		unsortedArray = nums2
		sort.Ints(nums1)
		sortedArray = nums1
	} else {
		unsortedArray = nums1
		sort.Ints(nums2)
		sortedArray = nums2
	}

	dict, res := map[int]bool{}, []int{}
	for _, v := range unsortedArray {
		if dict[v] == false {
			if helper349(sortedArray, v) {
				dict[v] = true
				res = append(res, v)
			}
		}
	}

	return res
}

// Binary search
func helper349(nums []int, target int) bool {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			return true
		}
	}
	return false
}
