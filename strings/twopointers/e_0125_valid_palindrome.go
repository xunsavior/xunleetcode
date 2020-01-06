package twopointers

import "unicode"

/*
Company: Facebook(41), Microsoft(9), Yandex(5), Apple(4), Google(2), Amazon(2)
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:
Input: "A man, a plan, a canal: Panama"
Output: true

Example 2:
Input: "race a car"
Output: false
*/

func isPalindrome125(s string) bool {
	n := len(s)
	if s == "" {
		return true
	}
	l, r := 0, n-1
	for l < r {
		lv, rv := rune(s[l]), rune(s[r])
		if !unicode.IsNumber(lv) && !unicode.IsLetter(lv) {
			l++
			continue
		}
		if !unicode.IsNumber(rv) && !unicode.IsLetter(rv) {
			r--
			continue
		}
		if unicode.ToLower(lv) != unicode.ToLower(rv) {
			return false
		}
		l++
		r--
	}
	return true
}
