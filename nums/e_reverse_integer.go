package nums

import (
	"math"
	"strconv"
	"strings"
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
func reverse(x int) int {
	if x > int(math.Pow(2, 31)-1) || x < -int(math.Pow(2, 31)) {
		return 0
	}
	s := strconv.Itoa(x)
	var actuallyLen int
	if actuallyLen = len(s) - 1; x >= 0 {
		actuallyLen = len(s)
	}
	sr := make([]string, actuallyLen, actuallyLen)
	var endpoint int
	if endpoint = 1; x > 0 {
		endpoint = 0
	}
	for i := len(s) - 1; i >= endpoint; i-- {
		sr = append(sr, string(s[i]))
	}
	sOutput := strings.Join(sr, "")
	out, err := strconv.Atoi(sOutput)
	if err == nil {
		if x > 0 && out <= int(math.Pow(2, 31)-1) {
			return out
		} else if x < 0 && -out >= -int(math.Pow(2, 31)) {
			return -out
		}
	}
	return 0
}
