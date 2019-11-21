package nums

import (
	"log"
	"math"
	"strconv"
)

/*
	Given a 32-bit signed integer, reverse digits of an integer.
	Example 1:
	Input: 123 Output: 321
	Example 2:
	Input: -123 Output: -321
	Example 3:
	Input: 120 Output: 21

	Note:
	Assume we are dealing with an environment which could only store integers within
	the 32-bit signed integer range: [−2^31,  2^31 − 1].
	For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

// TestReverseNumber ...
func TestReverseNumber() {
	log.Println(reverse(123))
	log.Println(reverse(-123))
	log.Println(reverse(120))
	log.Println(reverse(1534236469))
	log.Println(reverse(9646324351))
}

func reverse(x int) int {
	//case: the num is out of range
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	// convert num to string
	s := strconv.Itoa(x)
	initStr := ""
	runes := []rune(s)
	// check if 1st value is "-"
	var startIndex int
	if string(runes[startIndex]) == "-" {
		startIndex = 1
		initStr += "-"
	}
	//reverse and copy to a new string
	for i := len(runes) - 1; i >= startIndex; i-- {
		initStr += string(runes[i])
	}

	if num, err := strconv.Atoi(initStr); err == nil {
		return num
	}

	return 0
}
