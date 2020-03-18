package backtracking

import "strings"

/*
Company: Microsoft(15), Amazon(10), Facebook(10), Bloomberg(6), Adobe(5)

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

func generateParenthesis22(n int) []string {
	if n == 0 {
		return nil
	}
	res := []string{}
	left, right := strings.Repeat("(", n), strings.Repeat(")", n)
	helper22(left, right, []byte{}, &res)
	return res
}

func helper22(left, right string, newStr []byte, res *[]string) {
	if len(left) == 0 && len(right) == 0 {
		*res = append(*res, string(newStr))
		return
	}
	// if left-half ( is 0, then we just append right-half )
	if len(left) == 0 {
		appendedStr := append(newStr, right[0])
		helper22(left, right[1:], appendedStr, res)
		return
	}

	if len(left) == len(right) {
		// if left-half == right-half, remember we always append left half bracket first in order to form a valid parenthesis
		appendedStr := append(newStr, left[0])
		helper22(left[1:], right, appendedStr, res)
	} else {
		// we can either append "(" or ")"
		appendedStr1 := append(newStr, left[0])
		helper22(left[1:], right, appendedStr1, res)
		appendedStr2 := append(newStr, right[0])
		helper22(left, right[1:], appendedStr2, res)
	}

}
