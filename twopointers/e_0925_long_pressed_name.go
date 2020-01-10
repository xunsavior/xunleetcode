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
	sname, styped := len(name), len(typed)
	if sname == 0 || styped == 0 {
		return false
	}
	if name[0] != typed[0] {
		return false
	}
	i, j := 1, 1
	isLastCharFound := false
	for i < sname && j < styped {
		cname, ctyped := name[i], typed[j]
		if ctyped == cname {
			if i == sname-1 {
				isLastCharFound = true
			}
			i++
			j++
		} else {
			if ctyped == typed[j-1] {
				j++
				continue
			}
			return false
		}
	}
	return styped > sname && isLastCharFound
}
