package twopointers

/*
Company: Microsoft(5), Apple(4), Facebook(2), Google(2)

Write a function that reverses a string. The input string is given as an array of characters char[].

Do not allocate extra space for another array, you must do this by modifying the input array IN-PLACE
with O(1) extra memory.

You may assume all the characters consist of printable ascii characters.

Example 1:
Input: ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]

Example 2:
Input: ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
*/

func reverseString344(s []byte) {
	n := len(s)
	if n == 0 {
		return
	}
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
}
