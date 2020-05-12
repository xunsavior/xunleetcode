package hashtable

/*
Company: Amazon(2), Apple(2); Google(4); Adobe(6), Yahoo(3)

You're given strings J representing the types of stones that are jewels,
and S representing the stones you have.  Each character in S is a type
of stone you have.  You want to know how many of the stones you have
are also jewels.

The letters in J are guaranteed distinct, and all characters in J and S are letters.
Letters are case sensitive, so "a" is considered a different type of stone from "A".

Example 1:
Input: J = "aA", S = "aAAbbbb"
Output: 3

Example 2:
Input: J = "z", S = "ZZ"
Output: 0

Note:
S and J will consist of letters and have length at most 50.
The characters in J are distinct.
*/

func numJewelsInStones771(J string, S string) int {
	res, dict := 0, make([]bool, 52)
	for i := range J {
		if J[i] <= 90 && J[i] >= 65 {
			dict[J[i]-65] = true
		} else {
			dict[J[i]-71] = true
		}
	}

	for i := range S {
		if S[i] <= 90 && S[i] >= 65 {
			if dict[S[i]-65] {
				res++
			}
		} else {
			if dict[S[i]-71] {
				res++
			}
		}
	}

	return res
}
