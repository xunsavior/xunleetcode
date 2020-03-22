package backtracking

import "strconv"

/*
Company: nil; Amazon(4), Microsoft(2); Bloomberg(2), Google(2)

The set [1,2,3,...,n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:
"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.

Note:
Given n will be between 1 and 9 inclusive.
Given k will be between 1 and n! inclusive.

Example 1:
Input: n = 3, k = 3
Output: "213"

Example 2:
Input: n = 4, k = 9
Output: "2314"
*/

// Time: O(n^2)
// Space: O(n)
func getPermutation60(n int, k int) string {
	factorials, nums := make([]int, n), []int{1}
	factorials[0] = 1

	for i := 1; i < n; i++ {
		// generate factorial system bases 0!, 1!, ..., (n - 1)!
		factorials[i] = factorials[i-1] * i
		// generate nums 1, 2, ..., n
		nums = append(nums, i+1)
	}

	// fit k int the interval 0 ... (n!-1)
	fitK := k - 1

	// compute factorial representation of k
	str := ""
	for i := n - 1; i > -1; i-- {
		index := fitK / factorials[i]
		fitK -= index * factorials[i]
		str += strconv.Itoa(nums[index])
		nums = append(nums[:index], nums[index+1:]...) // remove the added number to avoid replication
	}

	return str
}

// Backtracking, not suggested
func getPermutation60WithBackTracking(n int, k int) string {
	res := []string{}
	initial := 1
	for i := n - 1; i > 0; i-- {
		initial *= i
	}

	start, fRes := 0, float64(k)/float64(initial)
	if fRes != float64(int(fRes)) {
		start = int(fRes) + 1
	} else {
		start = int(fRes)
	}

	remains := k % initial
	if remains == 0 {
		remains = initial
	}

	nums := make([]string, 0, n)
	nums = append(nums, strconv.Itoa(start))
	for i := 1; i <= n; i++ {
		if start != i {
			nums = append(nums, strconv.Itoa(i))
		}
	}

	helper60("", nums, &res, remains)
	return res[remains-1]
}

func helper60(str string, nums []string, res *[]string, k int) {
	if len(nums) == 0 {
		*res = append(*res, str)
		return
	}

	if len(*res) == k {
		return
	}

	for i := range nums {
		trimedNums := make([]string, 0, len(nums)-1)
		for j := 0; j < len(nums); j++ {
			if i != j {
				trimedNums = append(trimedNums, nums[j])
			}
		}
		helper60(str+nums[i], trimedNums, res, k)
	}
}
