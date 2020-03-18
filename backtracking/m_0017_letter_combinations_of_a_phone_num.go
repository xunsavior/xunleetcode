package backtracking

/*
Company: Amazon(18), Microsoft(8), Google(8), Facebook(5), Apple(4), Salesforce(4)

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example:
Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

Note:
Although the above answer is in lexicographical order, your answer could be in any order you want.
*/

var dic = map[int]string{
	50: "abc",
	51: "def",
	52: "ghi",
	53: "jkl",
	54: "mno",
	55: "pqrs",
	56: "tuv",
	57: "wxyz",
}

func letterCombinations17(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	res := []string{}
	helper17(0, []byte{}, digits, &res)
	return res
}

func helper17(index int, newStr []byte, digits string, res *[]string) {
	if index == len(digits) {
		*res = append(*res, string(newStr))
		return
	}
	charInt := int(digits[index])
	str := dic[charInt]
	for i := range str {
		appendedStr := append(newStr, str[i])
		helper17(index+1, appendedStr, digits, res)
	}
}
