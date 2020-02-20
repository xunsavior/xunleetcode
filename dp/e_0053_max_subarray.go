package dp

import "math"

/*
Company: Amazon(25), Microsoft(16), Google(16), LinkedIn(12), Apple(12)

Given an integer array nums, find the contiguous subarray (containing at least one number)
which has the largest sum and return its sum.

Example:
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Follow up:
If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

// Kadane’s Algorithm
// dp[i] = max(dp[i-1], 0) + arr[i]
func maxSubArray53DPSolution(nums []int) int {
	n := len(nums)
	maxSum := nums[0]

	for i := 1; i < n; i++ {
		// we add the current element to the maximum of 0 and previous sum
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		maxSum = int(math.Max(float64(nums[i]), float64(maxSum)))
	}

	return maxSum
}

func maxSubArray53DWithoutChangingOriginalArray(nums []int) int {
	maxSum, sum := math.MinInt32, 0

	for _, v := range nums {
		sum += v
		maxSum = int(math.Max(float64(maxSum), float64(sum)))
		// very important
		if sum < 0 {
			sum = 0
		}
	}

	return maxSum
}

// split nums into left, cross and right, then return sum of each group
func maxSubArray53DCSolution(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	mid := (len(nums) - 1) / 2

	leftSum := maxSubArray53DCSolution(nums[0 : mid+1])
	rightSum := maxSubArray53DCSolution(nums[mid+1:])
	crossSum := helper53(nums, mid)

	return int(math.Max(float64(crossSum), math.Max(float64(leftSum), float64(rightSum))))
}

func helper53(nums []int, mid int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	leftSubSum, rightSubSum, currSum := math.MinInt32, math.MinInt32, 0

	// important: from mid -> 0, as the subarray is consecutive
	for i := 0; i <= mid; i++ {
		currSum += nums[i]
		leftSubSum = int(math.Max(float64(leftSubSum), float64(currSum)))
	}

	currSum = 0
	for i := mid + 1; i < len(nums); i++ {
		currSum += nums[i]
		rightSubSum = int(math.Max(float64(rightSubSum), float64(currSum)))
	}

	return leftSubSum + rightSubSum
}
