package dp

import "math"

/*
Company: Amazon(122), Microsoft(29), Adobe(13), Facebook(10), Bloomberg(9), Google(8)

Given a string s, find the longest palindromic substring in s.
You may assume that the maximum length of s is 1000.

Example 1:
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example 2:
Input: "cbbd"
Output: "bb"
*/

/*
To improve over the brute force solution, we first observe how we can avoid unnecessary
re-computation while validating palindromes. Consider the case "ababa". If we already
knew that "bab" is a palindrome, it is obvious that "ababa" must be a palindrome since
the two left and right end letters are the same.
*/
func longestPalindrome5(s string) string {
	if len(s) == 0 {
		return ""
	}

	l, r := 0, 0

	for i := range s {
		len1, len2 := helper5(s, i, i), helper5(s, i, i+1)
		maxLen := int(math.Max(float64(len1), float64(len2)))
		if maxLen > r-l {
			l, r = i-(maxLen-1)/2, i+maxLen/2
		}
	}

	return s[l : r+1]
}

func helper5(s string, left, right int) int {
	if len(s) == 0 || left > right {
		return 0
	}

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1
}
