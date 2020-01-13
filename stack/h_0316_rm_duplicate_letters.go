package stack

/*
Company: null; Google(2), Nutanix(2), Amazon(2)

Given a string which contains only lowercase letters, remove duplicate
letters so that every letter appears once and only once. You must make
sure your result is the smallest in lexicographical order among all
possible results.

Example 1:
Input: "bcabc"
Output: "abc"

Example 2:
Input: "cbacdcbc"
Output: "acdb"
*/

func removeDuplicateLetters316(s string) string {
	dict, dictStack := make(map[int]int, 26), make(map[int]bool)
	stack := make([]int, 0)
	res := ""
	for _, v := range s {
		dict[int(v)]++
	}

	for _, v := range s {
		ascii := int(v)
		dict[ascii]--
		isLarger := true
		for len(stack) != 0 && ascii < stack[len(stack)-1] && dict[stack[len(stack)-1]] > 0 {
			// important: we need to check if the val has already been existed in the stack
			if dictStack[ascii] {
				break
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			dictStack[pop] = false
			isLarger = false
		}
		if isLarger && dictStack[ascii] {
			continue
		}
		stack = append(stack, ascii)
		dictStack[ascii] = true
	}

	for _, v := range stack {
		res += string(v)
	}

	return res
}
