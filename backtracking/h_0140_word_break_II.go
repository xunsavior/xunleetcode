package backtracking

import "strings"

/*
Company: Amazon(25), Facebook(12), Microsoft(5), Bloomberg(4), Google(3)

Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, add spaces in s
to construct a sentence where each word is a valid dictionary word. Return all such possible sentences.

Note:
The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.

Example 1:
Input:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
Output:
[
  "cats and dog",
  "cat sand dog"
]

Example 2:
Input:
s = "pineapplepenapple"
wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
Output:
[
  "pine apple pen apple",
  "pineapple pen apple",
  "pine applepen apple"
]
Explanation: Note that you are allowed to reuse a dictionary word.

Example 3:
Input:
s = "catsandog"
wordDict = ["cats", "dog", "sand", "and", "cat"]
Output:
[]
*/

func wordBreak140(s string, wordDict []string) []string {
	cache := make(map[string][]string)
	return helper140(s, cache, wordDict)
}

/*
  f(catsanddog) = cat + f(sanddog)
  f(sanddog) = sand + f(dog)
  f(dog) = dog + []string{}

  f(catsanddog) = cats + f(anddog)
  f(anddog) = and + f(dog)
  f(dog) = dog + []string{}

*/
// Time: O(n^3) Space: O(n^3)
func helper140(s string, cache map[string][]string, wordDict []string) []string {
	if res, ok := cache[s]; ok {
		return res
	}

	res := []string{}
	if len(s) == 0 {
		res = append(res, s) // important: used to add the last word
		return res
	}

	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			rightList := helper140(s[len(word):], cache, wordDict)
			for _, rightVal := range rightList {
				space := " "
				if len(rightVal) == 0 {
					space = ""
				}
				res = append(res, word+space+rightVal)
			}
		}
	}

	cache[s] = res
	return res
}
