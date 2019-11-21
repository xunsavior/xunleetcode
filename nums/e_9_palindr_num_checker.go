package nums

import (
	"fmt"
	"strconv"
)

/*
	Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

	Example 1:
	Input: 121
	Output: true

	Example 2:
	Input: -121
	Output: false
	Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

	Example 3:
	Input: 10
	Output: false
	Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

*/
func isPalindrome(x int) bool {
	stringx := strconv.Itoa(x)
	runes := []rune(stringx)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

//RunIsPalindrome ...
func RunIsPalindrome() {
	r1 := isPalindrome(121)
	r2 := isPalindrome(-121)
	r3 := isPalindrome(10)
	fmt.Printf("Input 121 Output %t\n", r1)
	fmt.Printf("Input -121 Output %t\n", r2)
	fmt.Printf("Input 10 Output %t\n", r3)
}
