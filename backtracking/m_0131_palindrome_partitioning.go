package backtracking

/*
Company: Uber(2), Yahoon(2), Qualtrics(2); Amazon(3), Adobe(3)

Given a string s, partition s such that every substring of the partition is a palindrome.

Return all possible palindrome partitioning of s.

Example:
Input: "aab"
Output:
[
  ["aa","b"],
  ["a","a","b"]
]
*/

func partition131(s string) [][]string {
	res := [][]string{}
	helper131(0, []string{}, s, &res)
	return res
}

func helper131(index int, tmp []string, s string, res *[][]string) {
	if index == len(s) {
		*res = append(*res, tmp)
		return
	}

	for i := index; i <= len(s); i++ {
		if isValid(s[index:i]) {
			t := append(tmp, s[index:i])
			helper131(i, t, s, res)
		}
	}

}

func isValid(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] == s[j] {
			return false
		}
	}
	return true
}
