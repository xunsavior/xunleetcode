package hashtable

/*
Company: Amazon(4), Microsoft(3), Google(2); Apple(2)

You are given a string, s, and a list of words, words, that are
all of the same length. Find all starting indices of substring(s)
in s that is a concatenation of each word in words exactly once
and without any intervening characters.

Example 1:
Input:
  s = "barfoothefoobarman",
  words = ["foo","bar"]
Output: [0,9]
Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar" respectively.
The output order does not matter, returning [9,0] is fine too.

Example 2:
Input:
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
Output: []
*/

// O(n^2)
func findSubstring30(s string, words []string) []int {
	if s == "" || len(words) == 0 {
		return nil
	}
	res := []int{}
	lword, lwords := len(words[0]), len(words)
	dict := make(map[string]int)
	for _, v := range words {
		dict[v]++
	}

	for i := 0; i <= len(s)-lwords*lword; i++ {
		copy := make(map[string]int)
		for k, v := range dict {
			copy[k] = v
		}
		k, j := lwords, i
		for k > 0 {
			str := s[j : j+lword]
			if copy[str] < 1 {
				break
			}
			copy[str]--
			k--
			j += lword
			if k == 0 {
				res = append(res, i)
			}
		}
	}
	return res
}
