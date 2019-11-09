package strs

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
(you may want to display this pattern in a fixed font for better legibility)
P   A   H   N
A P L S I I G
Y   I   R

Example 1:
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"

Example 2:
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I
*/
const (
	MoveDown = -1
	MoveUp   = 1
)

// Convert ...
func Convert(s string, numRows int) string {
	// if s only contains one char
	if numRows == 1 {
		return s
	}

	// init direction and row number
	dir, rowNo := MoveDown, 0
	runes := []rune(s)

	// use hash map with key: row index and value: string of each row
	dic := map[int]string{}
	for _, v := range runes {
		if dir == MoveDown {
			if rowNo <= numRows-1 {
				dic[rowNo] = dic[rowNo] + string(v)
				if rowNo == numRows-1 {
					rowNo--
					dir = MoveUp
					continue
				}
				rowNo++
			}
		} else {
			if rowNo >= 0 {
				dic[rowNo] = dic[rowNo] + string(v)
				if rowNo == 0 {
					rowNo++
					dir = MoveDown
					continue
				}
				rowNo--
			}
		}
	}

	//pending string by traversing the map
	var result string
	for i := 0; i < numRows; i++ {
		result += dic[i]
	}

	return result
}
