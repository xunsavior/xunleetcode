package strs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
	Given a string, find the length of the longest substring without repeating characters.

	Example 1:
	Input: "abcabcbb"
	Output: 3
	Explanation: The answer is "abc", with the length of 3.

	Example 2:
	Input: "bbbbb"
	Output: 1
	Explanation: The answer is "b", with the length of 1.

	Example 3:
	Input: "pwwkew"
	Output: 3

	Example 4:
	Input: " "
	Output: 1
	Explanation: The answer is "wke", with the length of 3.
    Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
	//create empty string to store substr
	substr := ""
	maxLen := 0

	for _, v := range s {
		strValue := string(v)
		repeatedIndex := strings.Index(substr, strValue) //get the first repeated index in substr
		substr += strValue                               // append the value to substr

		// if the repeated index is found, we discard all the elements equal or less the index.
		if repeatedIndex != -1 {
			substr = substr[repeatedIndex+1:]
		}

		// update the max size
		if size := utf8.RuneCountInString(substr); size >= maxLen {
			maxLen = size
		}
	}

	return maxLen
}

// RunLongestNonRepatedSubstr ...
func RunLongestNonRepatedSubstr() {
	result1 := lengthOfLongestSubstring("abcabcbb")
	result2 := lengthOfLongestSubstring("bbbbb")
	result3 := lengthOfLongestSubstring("pwwkew")
	result4 := lengthOfLongestSubstring(" ")
	fmt.Printf("Input abcabcbb Output %d\n", result1)
	fmt.Printf("Input bbbbb Output %d\n", result2)
	fmt.Printf("Input pwwkew Output %d\n", result3)
	fmt.Printf("Input \" \" Output %d\n", result4)
}
