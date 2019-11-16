package leetcodetest

import (
	nums "leetcode/nums"
	"testing"
)

// Test Int to Roman
type intToRoman struct {
	num   int
	roman string
}

var intToRomanCases = []*intToRoman{
	&intToRoman{
		num:   3,
		roman: "III",
	},
	&intToRoman{
		num:   4,
		roman: "IV",
	},
	&intToRoman{
		num:   58,
		roman: "LVIII",
	},
	&intToRoman{
		num:   1994,
		roman: "MCMXCIV",
	},
	&intToRoman{
		num:   10,
		roman: "X",
	},
}

func TestMIntToRoman(t *testing.T) {
	for _, v := range intToRomanCases {
		equals(t, v.roman, nums.IntToRoman(v.num))
	}
}
