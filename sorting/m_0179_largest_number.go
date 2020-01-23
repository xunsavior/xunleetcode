package sorting

import (
	"sort"
	"strconv"
	"strings"
)

/*
Company: Microsoft(4), Amazon(4); Salesforce(2); Facebook(2)
Given a list of non negative integers, arrange them such that they form the largest number.

Example 1:
Input: [10,2]
Output: "210"

Example 2:
Input: [3,30,34,5,9]
Output: "9534330"
Note: The result may be very large, so you need to return a string instead of an integer.
*/

func largestNumber179(nums []int) string {
	strnums := make([]string, 0)
	for _, v := range nums {
		strnums = append(strnums, strconv.Itoa(v))
	}
	sort.Slice(strnums, func(i, j int) bool {
		return (strnums[i] + strnums[j]) > (strnums[j] + strnums[i])
	})
	numStr := strings.Join(strnums, "")
	if num, _ := strconv.Atoi(numStr); num == 0 {
		return "0"
	}
	return numStr
}
