package hashtable

import (
	"sort"

	"github.com/xunsavior/xunleetcode/utils"
)

/*
Company: Amazon(23), Microsoft(19), Uber(5), Apple(4), Google(4), Facebook(3)
Given an array of strings, group anagrams together.

Example:
Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]

Note:
All inputs will be in lowercase.
The order of your output does not matter.
*/

func groupAnagrams49(strs []string) [][]string {
	res, dict := make([][]string, 0), make(map[string]int)
	for _, v := range strs {
		runes := utils.StringToRunes(v)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		str := string(runes)
		if index, ok := dict[str]; ok {
			res[index] = append(res[index], v)
		} else {
			words := []string{v}
			dict[str] = len(res)
			res = append(res, words)
		}
	}
	return res
}
