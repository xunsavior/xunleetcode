package bs

import "sort"

/*
Company: Facebook(9), Amazon(5), LinkedIn(2), Microsoft(2), Yandex(2)

Given two arrays, write a function to compute their intersection.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]

Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]

Note:
Each element in the result should appear as many times as it shows in both arrays.
The result can be in any order.

Follow up:
What if the given array is already sorted? How would you optimize your algorithm?
What if nums1's size is small compared to nums2's size? Which algorithm is better?
What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
*/

// Time: O(M*logN)
// Space: O(M+N)
func intersect350(nums1 []int, nums2 []int) []int {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 || n2 == 0 {
		return nil
	}

	var sortedList, unsortedList []int
	if n1 < n2 {
		unsortedList = nums2
		sort.Ints(nums1)
		sortedList = nums1
	} else {
		unsortedList = nums1
		sort.Ints(nums2)
		sortedList = nums2
	}

	dict, res := make(map[int]int), []int{}
	for _, v := range sortedList {
		dict[v]++
	}

	for _, v := range unsortedList {
		if dict[v] > 0 {
			if helper350(sortedList, v) {
				dict[v]--
				res = append(res, v)
			}
		}
	}

	return res
}

func helper350(nums []int, target int) bool {
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
