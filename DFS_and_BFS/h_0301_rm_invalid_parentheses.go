package dfsandbfs

/*
Company: Facebook(58), Uber(3), Amazon(2), Bloomberg(3), Postmates(3)

Remove the MINIMUM number of invalid parentheses in order to make the input string valid. Return all possible results.

Note: The input string may contain letters other than the parentheses ( and ).

Example 1:
Input: "()())()"
Output: ["()()()", "(())()"]

Example 2:
Input: "(a)())()"
Output: ["(a)()()", "(a())()"]

Example 3:
Input: ")("
Output: [""]
*/

// Time: O(2^(l+r))
func removeInvalidParentheses301(s string) []string {
	res := []string{}
	// total num of "(" and ")" needs to be removed
	l, r := 0, 0
	for i := range s {
		if s[i] == 40 {
			l++
		}
		if l == 0 && s[i] == 41 {
			// if l == 0 and current char is ")", that means this ")" is redundant
			if s[i] == 41 {
				r++
			}
		} else {
			// if l > 0 and current char is ")", that means we need to reduce l by 1
			if s[i] == 41 {
				l--
			}
		}
	}
	helper301(s, &res, 0, l, r) // the deleted char could be "(" or ")"
	return res
}

// check if the parentheses is valid or not
func isValid301(s string) bool {
	count := 0
	for i := range s {
		if s[i] == 40 {
			count++
		}

		if s[i] == 41 {
			count--
		}

		if count < 0 {
			return false
		}
	}
	return count == 0
}

// DFS
func helper301(s string, res *[]string, start int, l, r int) {
	if l == 0 && r == 0 {
		if isValid301(s) {
			*res = append(*res, s)
		}
		return
	}

	for i := start; i < len(s); i++ {
		// we only remove the first parenthes if there are consecutive ones to avoid duplication
		if i != start && s[i] == s[i-1] {
			continue
		}

		if s[i] == 40 || s[i] == 41 {
			str := s[:i] + s[i+1:]
			// we must remove ")" first then "("
			if r > 0 {
				helper301(str, res, i, l, r-1)
			} else if l > 0 {
				helper301(str, res, i, l-1, r)
			}
		}
	}
}
