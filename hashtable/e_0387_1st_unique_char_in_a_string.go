package hashtable

/*
Top interviewed question
Company: Amazon(13), Goldman Sachs(13), Bloomberg(7), Microsoft(5), Zillow(2)

Given a string, find the first non-repeating character in it and return it's index.
If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
Note: You may assume the string contain only lowercase letters.
*/

func firstUniqChar387(s string) int {
	res := -1
	dict := make([]int, 26)
	for i := range s {
		index := s[i] - 97
		dict[index]++
	}
	for i := range s {
		index := s[i] - 97
		if dict[index] == 1 {
			res = i
			break
		}
	}
	return res
}
