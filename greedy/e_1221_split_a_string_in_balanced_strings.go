package greedy

/*
Company: Apple(6); nil; nil

Balanced strings are those who have equal quantity of 'L' and 'R' characters.
Given a balanced string s split it in the maximum amount of balanced strings.

Return the maximum amount of splitted balanced strings.

Example 1:
Input: s = "RLRRLLRLRL"
Output: 4
Explanation: s can be split into "RL", "RRLL", "RL", "RL", each substring contains same number of 'L' and 'R'.

Example 2:
Input: s = "RLLLLRRRLR"
Output: 3
Explanation: s can be split into "RL", "LLLRRR", "LR", each substring contains same number of 'L' and 'R'.

Example 3:
Input: s = "LLLLRRRR"
Output: 1
Explanation: s can be split into "LLLLRRRR".

Example 4:
Input: s = "RLRRRLLRLL"
Output: 2
Explanation: s can be split into "RL", "RRRLLRLL", since each substring contains an equal number of 'L' and 'R'

Constraints:
1 <= s.length <= 1000
s[i] = 'L' or 'R'
*/

// time: O(N)
// space: O(1)
func balancedStringSplit1221(s string) int {
	numL, numR, res := 0, 0, 0
	L := []byte("L")[0]
	for i := range s {
		if s[i] == L {
			numL++
		} else {
			numR++
		}
		if numL > 0 && numR > 0 && numL == numR {
			numL, numR = 0, 0
			res++
		}
	}
	return res
}
