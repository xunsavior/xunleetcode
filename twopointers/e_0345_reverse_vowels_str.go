package twopointers

/*
Company: Amazon(2)
Write a function that takes a string as input and reverse only the vowels of a string.

Example 1:
Input: "hello"
Output: "holle"

Example 2:
Input: "leetcode"
Output: "leotcede"

Note:
The vowels does not include the letter "y".
*/

func reverseVowels345(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	chars := []byte(s)
	l, r := 0, n-1
	vowelsMap := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
		"A": true,
		"E": true,
		"I": true,
		"O": true,
		"U": true,
	}
	for l < r {
		lv, rv := string(chars[l]), string(chars[r])
		if !vowelsMap[lv] {
			l++
			continue
		}
		if !vowelsMap[rv] {
			r--
			continue
		}
		chars[l], chars[r] = chars[r], chars[l]
		l++
		r--
	}
	return string(chars)
}
