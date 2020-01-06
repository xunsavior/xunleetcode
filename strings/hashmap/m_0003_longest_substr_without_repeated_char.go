package twopointers

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
	n := len(s)
	if n < 1 {
		return n
	}
	start, res := -1, 0
	charMap := make(map[string]int)
	for i := range s {
		str := string(s[i])
		index, ok := charMap[str]
		if ok && index > start {
			start, charMap[str] = charMap[str], i
		} else {
			charMap[str] = i
			if i-start > res {
				res = i - start
			}
		}
	}
	return res
}
