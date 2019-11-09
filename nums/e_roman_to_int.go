package nums

import (
	"fmt"
)

/*
	Example 1:
	Input: "III"
	Output: 3

	Example 2:
	Input: "IV"
	Output: 4

	Example 3:
	Input: "IX"
	Output: 9

	Example 4:
	Input: "LVIII"
	Output: 58
	Explanation: L = 50, V= 5, III = 3.

	Example 5:
	Input: "MCMXCIV"
	Output: 1994
	Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

	IV, IX
	XL, XC
	CD, CM
*/
const (
	I = 1
	V = 5
	X = 10
	L = 50
	C = 100
	D = 500
	M = 1000
)

var romanMap = map[string]int{
	"I": I,
	"V": V,
	"X": X,
	"L": L,
	"C": C,
	"D": D,
	"M": M,
}

func romanToInt(s string) int {
	sum := 0
	for i, v := range s {
		c := string(v)     // get roman char
		val := romanMap[c] // get value of the above roman char

		// check if we need to change the sign of the value of roman
		switch c {
		case "I":
			if i+1 < len(s) && (string(s[i+1]) == "V" || string(s[i+1]) == "X") {
				val = -val
			}
		case "X":
			if i+1 < len(s) && (string(s[i+1]) == "L" || string(s[i+1]) == "C") {
				val = -val
			}
		case "C":
			if i+1 < len(s) && (string(s[i+1]) == "D" || string(s[i+1]) == "M") {
				val = -val
			}
		}

		// sum up
		sum += val
	}

	return sum
}

//RumRomanToInt ...
func RumRomanToInt() {
	fmt.Println("III", romanToInt("III"))
	fmt.Println("IV", romanToInt("IV"))
	fmt.Println("IX", romanToInt("IX"))
	fmt.Println("LVIII", romanToInt("LVIII"))
	fmt.Println("MCMXCIV", romanToInt("MCMXCIV"))
}
