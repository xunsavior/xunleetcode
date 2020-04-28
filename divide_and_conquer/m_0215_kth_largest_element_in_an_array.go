package divideandconquer

/*
Company: Facebook(26), Amazon(16), LinkedIn(7), Google(5), Microsoft(4)

Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.
*/

// Time: O(N) to O(N^2) Space: O(1)
func findKthLargest215(nums []int, k int) int {
	start, end := 0, len(nums)-1
	pivotIndex := helper215(nums, start, end)
	for pivotIndex != len(nums)-k {
		if pivotIndex > len(nums)-k {
			end = pivotIndex - 1
		} else {
			start = pivotIndex + 1
		}
		pivotIndex = helper215(nums, start, end)
	}
	return nums[pivotIndex]
}

// partition
func helper215(nums []int, start, end int) int {
	pivot := start + (end-start)/2
	nums[pivot], nums[end] = nums[end], nums[pivot]

	i := start
	for j := start; j < end; j++ {
		if nums[j] <= nums[end] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	return i
}
