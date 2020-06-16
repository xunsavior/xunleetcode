package twopointers

/*
Company: Google(3)
Your friend is typing his name into a keyboard.  Sometimes, when typing a character c,
the key might get long pressed, and the character will be typed 1 or more times.

You examine the typed characters of the keyboard.  Return True if it is possible that
it was your friends name, with some characters (possibly none) being long pressed.

Example 1:
Input: name = "alex", typed = "aaleex"
Output: true
Explanation: 'a' and 'e' in 'alex' were long pressed.

Example 2:
Input: name = "saeed", typed = "ssaaedd"
Output: false
Explanation: 'e' must have been pressed twice, but it wasn't in the typed output.

Example 3:
Input: name = "leelee", typed = "lleeelee"
Output: true

Example 4:
Input: name = "laiden", typed = "laiden"
Output: true

Explanation: It's not necessary to long press any character.

Note:
name.length <= 1000
typed.length <= 1000
The characters of name and typed are lowercase letters.
*/

func isLongPressedName925(name string, typed string) bool {
	i, j := 0, 0
	for j < len(typed) {

		if i < len(name) && name[i] == typed[j] {
			i++
			j++
			continue
		}

		if j == 0 || typed[j] != typed[j-1] {
			return false
		}

		j++
	}

	return i == len(name)
}

/*
Test cases:
"alex"
"aaleex"
"saeed"
"ssaaedd"
"leelee"
"lleeelee"
"laiden"
"laiden"
"pyplrz"
"ppyypllr"
"vtkgn"
"vttkgnn"
*/
