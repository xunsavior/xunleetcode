package dp

/*
Company: Facebook(8), Amazon(5), Twitter(4), Google(2), Uber(2)

Given a string, your task is to count how many palindromic substrings in this string.

The substrings with different start indexes or end indexes are counted as different
substrings even they consist of same characters.

Example 1:
Input: "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".

Example 2:
Input: "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".


Note:
The input string length won't exceed 1000.
*/

// Time: O(n^2), Space: O(n^2)
func countSubstrings647(s string) int {
	n := len(s)
	res := 0
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = true
			} else if i+1 == j {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][i-1]
			}
			if dp[i][j] == true {
				res++
			}
		}
	}

	return res
}

// Time: O(n^2), Space: O(1)
func countSubstrings647ExpansionSolution(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	for i := range s {
		res += helper647(s, i, i)
		res += helper647(s, i, i+1)
	}
	return res
}

func helper647(s string, l, r int) int {
	count := 0

	for l >= 0 && r < len(s) && s[l] == s[r] {
		count++
		l--
		r++
	}

	return count
}
