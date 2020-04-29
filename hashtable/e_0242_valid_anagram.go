package hashtable

/*
Top interviewed questions
Company: Amazon(6), Oracle(4), Google(3), Facebook(2), Bloomberg(2)

Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false

Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?
*/

func isAnagram242(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cache := make([]int, 26)
	for i := range s {
		index := s[i] - 97
		cache[index]++
	}
	for i := range t {
		index := t[i] - 97
		if cache[index] == 0 {
			return false
		}
		cache[index]--
	}

	return true
}
