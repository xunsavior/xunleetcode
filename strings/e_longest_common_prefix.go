package strs

/*
	Write a function to find the longest common prefix string amongst an array of strings.
	If there is no common prefix, return an empty string "".

	Example 1:
	Input: ["flower","flow","flight"]
	Output: "fl"

	Example 2:
	Input: ["dog","racecar","car"]
	Output: ""

	Explanation: There is no common prefix among the input strings.
	All given inputs are in lowercase letters a-z.
*/
import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	// check if it is empty slice
	if len(strs) == 0 {
		return ""
	}

	// find a shortest string, and use the shortest string to compare
	var shortestStr string = strs[0]
	for _, str := range strs {
		if len(str) < len(shortestStr) {
			shortestStr = str
		}
	}

	// find 1st index of char different from shortestStr's prefix
	for i, c := range shortestStr {
		for _, str := range strs {
			if string(str[i]) != string(c) {
				return shortestStr[:i] // return the substr
			}
		}
	}

	return shortestStr
}

//TestLongestCommonPrefix ...
func TestLongestCommonPrefix() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{}))
}
