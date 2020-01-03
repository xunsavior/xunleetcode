package twopointers

import (
	"strings"
	"unicode/utf8"
)

/*
Company: Amazon(37), Microsoft(14), Cisco(13), Google(12), Facebook(11)
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

Explanation: The answer is "wke", with the length of 3.
Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring3(s string) int {
	substr := ""
	maxLen := 0

	for _, v := range s {
		strValue := string(v)
		repeatedIndex := strings.Index(substr, strValue)
		substr += strValue
		if repeatedIndex != -1 {
			substr = substr[repeatedIndex+1:]
		}
		if size := utf8.RuneCountInString(substr); size >= maxLen {
			maxLen = size
		}
	}

	return maxLen
}
