package strs

import (
	"log"
	"math"
	"strconv"
	"strings"
)

/*
The function first discards as many whitespace characters as necessary until the first non-whitespace character is found.
Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible,
and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number,
which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number,
or if no such sequence exists because either str is empty or it contains only whitespace characters,
no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Note:

Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store integers
within the 32-bit signed integer range: [−231,  231 − 1].
If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.

*/

var max = math.MaxInt32
var min = math.MinInt32

func myAtoi(str string) int {
	// remove space
	trimedStr := strings.TrimSpace(str)
	numString := ""

	// used to check if positive or negative
	sign := 1
	for i, v := range trimedStr {
		// check first char is "+" or "-"
		if i == 0 {
			if string(v) == "+" {
				continue
			} else if string(v) == "-" {
				sign = -1 // update sign to -1
				continue
			}
		}
		// check if it is numeric
		if v <= 57 && v >= 48 {
			numString += string(v)
		} else {
			break
		}
	}

	if numString == "" {
		return 0
	}

	if num, err := strconv.Atoi(numString); err == nil {
		//check if it is over int32
		if num > max || num < min {
			if sign == 1 {
				return max
			}
			return min
		}
		return sign * num
	}

	//check if it is over int64
	if sign == 1 {
		return max
	}

	return min
}

//TestMyAoti ...
func TestMyAoti() {
	i0 := "20000000000000000000"
	o0 := myAtoi("20000000000000000000")
	log.Printf("Input: %s output: %d", i0, o0)

	i1 := "   -42"
	o1 := myAtoi("   -42")
	log.Printf("Input: %s output: %d", i1, o1)

	i2 := "4193 with words"
	o2 := myAtoi("4193 with words")
	log.Printf("Input: %s output: %d", i2, o2)

	i3 := "words and 987"
	o3 := myAtoi("words and 987")
	log.Printf("Input: %s output: %d", i3, o3)

	i4 := "+1"
	o4 := myAtoi("+1")
	log.Printf("Input: %s output: %d", i4, o4)

	i5 := "-91283472332"
	o5 := myAtoi("-91283472332")
	log.Printf("Input: %s output: %d", i5, o5)

}
