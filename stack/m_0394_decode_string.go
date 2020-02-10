package stack

/*
Company: Bloomberg(18), Google(7), Facebook(6), Amazon(5), Apple(3)

Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square
brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; No extra white spaces, square brackets
are well-formed, etc.

Furthermore, you may assume that the original data does not contain any digits and that digits
are only for those repeat numbers, k. For example, there won't be input like 3a or 2[4].

Examples:
s = "3[a]2[bc]", return "aaabcbc".
s = "3[a2[c]]", return "accaccacc".
s = "2[abc]3[cd]ef", return "abcabccdcdcdef".
*/

func decodeString394(s string) string {
	counts, results := make([]int, 0), make([]string, 0)
	res := ""
	index := 0
	for index < len(s) {
		if int(s[index]) > 47 && int(s[index]) < 58 { // numeric
			count := 0
			for index < len(s) && int(s[index]) > 47 && int(s[index]) < 58 {
				count = 10*count + int(s[index]) - 48
				index++
			}
			counts = append(counts, count)
		} else if int(s[index]) == 91 { // "["
			results = append(results, res)
			res = ""
			index++
		} else if int(s[index]) == 93 { // "]"
			popCount, resultPop := counts[len(counts)-1], results[len(results)-1]
			counts, results = counts[:len(counts)-1], results[:len(results)-1]
			for i := 0; i < popCount; i++ {
				resultPop += res
			}
			res = resultPop
			index++
		} else {
			res += string(int(s[index]))
			index++
		}
	}
	return res
}
