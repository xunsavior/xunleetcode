package hashtable

import (
	"math"
	"strconv"
)

/*
Company: JPMorgan(111), Google(5), Apple(4), Amazon(3), Microsoft(3)

Write an algorithm to determine if a number is "happy".

A happy number is a number defined by the following process: Starting with any positive integer,
replace the number by the sum of the squares of its digits, and repeat the process until the
number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy numbers.

Example:
Input: 19
Output: true
Explanation:
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1
*/

func isHappy202(n int) bool {
	dict := make(map[int]bool)
	sum := n
	for dict[sum] == false {
		dict[sum] = true
		sum = 0
		for _, v := range strconv.Itoa(sum) {
			num, _ := strconv.Atoi(string(v))
			sum += int(math.Pow(float64(num), 2))
		}
		if sum == 1 {
			return true
		}
	}
	return false
}
