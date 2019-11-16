package leetcodetest

import (
	strs "leetcode/strings"
	"testing"
)

// Test ZigZag Conversion
type zigzag struct {
	s         string
	numRows   int
	expResult string
}

var zigzagCases = []*zigzag{
	&zigzag{
		s:         "PAYPALISHIRING",
		numRows:   3,
		expResult: "PAHNAPLSIIGYIR",
	},
	&zigzag{
		s:         "PAYPALISHIRING",
		numRows:   4,
		expResult: "PINALSIGYAHRPI",
	},
}

func TestMZigzag(t *testing.T) {
	for _, v := range zigzagCases {
		equals(t, v.expResult, strs.Convert(v.s, v.numRows))
	}
}

// Valid Parentheses
type validParenthesesInput struct {
	s         string
	expectRes bool
}

var validParenthesesTestCases = []*validParenthesesInput{
	&validParenthesesInput{
		s:         "()",
		expectRes: true,
	},
	&validParenthesesInput{
		s:         "()[]{}",
		expectRes: true,
	},
	&validParenthesesInput{
		s:         "(]",
		expectRes: false,
	},
	&validParenthesesInput{
		s:         "([)]",
		expectRes: false,
	},
	&validParenthesesInput{
		s:         "{[]}",
		expectRes: true,
	},
	&validParenthesesInput{
		s:         "(([]){})",
		expectRes: true,
	},
}

func TestValidParentheses(t *testing.T) {
	for _, v := range validParenthesesTestCases {
		equals(t, v.expectRes, strs.IsValid(v.s))
	}
}
