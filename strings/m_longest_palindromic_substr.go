package strs

import "fmt"

/*
	Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

	Example 1:
	Input: "babad"
	Output: "bab"
	Note: "aba" is also a valid answer.

	Example 2:
	Input: "cbbd"
	Output: "bb"
	cbbbc
	aagabcdcba
*/

func longestPalindrome(s string) string {
	// runes := []rune(s)
	// if len(runes) == 1 {
	// 	return string(runes[0])
	// }
	// for i, r := range runes {
	// 	if i == 0 {

	// 	}
	// }
	return ""
}

//RunLongestPalindrome ...
func RunLongestPalindrome() {
	result1 := longestPalindrome("babad")
	result2 := longestPalindrome("cbbd")
	fmt.Printf("Input babad Output %s\n", result1)
	fmt.Printf("Input cbbd Output %s\n", result2)
}
