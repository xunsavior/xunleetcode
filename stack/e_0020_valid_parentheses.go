package stack

/*
Company: Amazon(32), Microsoft(14), Facebook(12), Google(6), LinkedIn(5)
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:
Input: "()"
Output: true

Example 2:
Input: "()[]{}"
Output: true

Example 3:
Input: "(]"
Output: false

Example 4:
Input: "([)]"
Output: false

Example 5:
Input: "{[]}"
Output: true
*/

// IsValid ...
func isValid20(s string) bool {
	if s == "" {
		return true
	}
	bytes := []byte{}
	for i := range s {
		if string(s[i]) == "(" {
			bytes = append(bytes, s[i]+1)
		} else if string(s[i]) == "[" || string(s[i]) == "{" {
			bytes = append(bytes, s[i]+2)
		} else {
			n := len(bytes)
			if n == 0 {
				return false
			}
			pop := bytes[n-1]
			bytes = bytes[0 : n-1]
			if pop != s[i] {
				return false
			}
		}
	}
	return len(bytes) == 0
}
