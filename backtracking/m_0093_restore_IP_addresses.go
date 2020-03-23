package backtracking

import "strconv"

/*
Company: Amazon(7), Microsoft(6), Facebook(3), VMware(2)

Given a string containing only digits, restore it by returning all possible valid IP address combinations.

Example:
Input: "25525511135"
Output: ["255.255.11.135", "255.255.111.35"]
*/

func restoreIPAddresses93(s string) []string {
	if len(s) < 4 {
		return nil
	}
	res := []string{}
	helper93(0, "", s, &res)
	return res
}

func helper93(index int, element, remainedSuffix string, res *[]string) {
	remainedSlot := 4 - index

	if remainedSlot == 0 || len(remainedSuffix) == 0 {
		if remainedSlot == len(remainedSuffix) {
			*res = append(*res, element)
		}
		return
	}

	if remainedSlot > 0 && len(remainedSuffix)/remainedSlot > 3 {
		return
	}

	dot := "."
	if remainedSlot == 1 {
		dot = ""
	}

	if len(remainedSuffix) >= 1 {
		oneChar := element + remainedSuffix[0:1] + dot
		helper93(index+1, oneChar, remainedSuffix[1:], res)
	}

	if len(remainedSuffix) >= 2 {
		// important: we must ensure that the first element is not 0, e.g. 01
		if remainedSuffix[0] != 48 {
			twoChars := element + remainedSuffix[0:2] + dot
			helper93(index+1, twoChars, remainedSuffix[2:], res)
		}
	}

	if len(remainedSuffix) >= 3 {
		segment := remainedSuffix[0:3]
		if num, err := strconv.Atoi(segment); err == nil {
			// important: we must ensure that the number should less than 255 as well as the first element must not 0, e.g. 001 or 011
			if num <= 255 && remainedSuffix[0] != 48 {
				threeChars := element + segment + dot
				helper93(index+1, threeChars, remainedSuffix[3:], res)
			}
		}
	}
}
