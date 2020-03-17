package backtracking

/*
Company: nil; Microsoft(2), Amazon(2); Facebook

Given a string S, we can transform every letter individually to be lowercase or uppercase to create another string.
Return a list of all possible strings we could create.

Examples:
Input: S = "a1b2"
Output: ["a1b2", "a1B2", "A1b2", "A1B2"]

Input: S = "3z4"
Output: ["3z4", "3Z4"]

Input: S = "12345"
Output: ["12345"]

Note:
S will be a string with length between 1 and 12.
S will consist only of letters or digits.
*/

func letterCasePermutation784(S string) []string {
	res := []string{}
	helper784(S, 0, &res)
	return res
}

func helper784(S string, index int, res *[]string) {
	if index == len(S) {
		*res = append(*res, S)
		return
	}

	if S[index] <= 57 && S[index] >= 48 {
		helper784(S, index+1, res)
		return
	}

	helper784(S, index+1, res)

	bytes := []byte(S)
	if S[index] >= 97 {
		bytes[index] -= 32
	} else {
		bytes[index] += 32
	}
	helper784(string(bytes), index+1, res)
}
