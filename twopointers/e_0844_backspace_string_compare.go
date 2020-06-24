package twopointers

/*
Company: Google(8), Facebook(7), Amazon(5), Microsoft(3), Oracle(3)

Given two strings S and T, return if they are equal when both are typed into empty text editors.
# means a backspace character.

Note that after backspacing an empty text, the text will continue empty.
Example 1:
Input: S = "ab#c", T = "ad#c"
Output: true
Explanation: Both S and T become "ac".

Example 2:
Input: S = "ab##", T = "c#d#"
Output: true
Explanation: Both S and T become "".

Example 3:
Input: S = "a##c", T = "#a#c"
Output: true
Explanation: Both S and T become "c".

Example 4:
Input: S = "a#c", T = "b"
Output: false
Explanation: S becomes "c" while T becomes "b".

Note:
1 <= S.length <= 200
1 <= T.length <= 200
S and T only contain lowercase letters and '#' characters.

Follow up:
Can you solve it in O(N) time and O(1) space?
*/

// We iterate from end to start
func backspaceCompare844(S string, T string) bool {
	i, j := len(S)-1, len(T)-1
	num1, num2 := 0, 0 // the number of unhandled # respectively
	for i >= 0 || j >= 0 {
		s, t := 0, 0 // 0 means there isn't any chars left

		if i >= 0 {
			s = int(S[i])
		}

		if j >= 0 {
			t = int(T[j])
		}

		if s == 35 {
			num1++
			i--
			continue
		}

		if t == 35 {
			num2++
			j--
			continue
		}

		if num1 > 0 {
			num1--
			i--
			continue
		}

		if num2 > 0 {
			num2--
			j--
			continue
		}
		// we must ensure that both num1 and num2 are 0 before making any comparison
		if s == t {
			i--
			j--
			continue
		}
		return false
	}
	return true
}

/*
Test Data
"ab#c"
"ad#c"
"ab##"
"c#d#"
"a##c"
"#a#c"
"a#c"
"b"
"bxj##tw"
"bxj###tw"
*/
