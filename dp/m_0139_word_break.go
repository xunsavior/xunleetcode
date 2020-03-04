package dp

/*
Company: Facebook(36), Amazon(24), Bloomberg(12), Microsoft(8), Google(7)

Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
determine if s can be segmented into a space-separated sequence of one or more dictionary
words.

Note:
The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.

Example 1:
Input: s = "leetcode", wordDict = ["leet", "code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:
Input: s = "applepenapple", wordDict = ["apple", "pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
             Note that you are allowed to reuse a dictionary word.
Example 3:
Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
Output: false
*/

func wordBreak139(s string, wordDict []string) bool {
	dic := make(map[string]bool)
	for _, v := range wordDict {
		dic[v] = true
	}
	cache := make([]*bool, len(s))
	return helper138(s, 0, dic, cache)
}

func helper138(s string, start int, wordDict map[string]bool, cache []*bool) bool {
	if start == len(s) {
		return true
	}

	if res := cache[start]; res != nil {
		return *res
	}
	res := false
	for i := start + 1; i <= len(s); i++ {
		text := s[start:i]

		if isExisted := wordDict[text]; !isExisted {
			continue
		}

		if isFollowingExisted := helper138(s, i, wordDict, cache); !isFollowingExisted {
			continue
		}
		res = true
		cache[start] = &res
		return res
	}

	res = false
	cache[start] = &res
	return res
}
