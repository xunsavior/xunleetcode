package backtracking

/*
Company: Amazon(16), Intuit(13), Facebook(9), Microsoft(5), Snapchat(5)

Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent"
cells are those horizontally or vertically neighboring. The same letter cell may not
be used more than once.

Example:
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]
Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.

Constraints:
board and word consists only of lowercase and uppercase English letters.
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3
*/

// Time O(m*n*4^l) l is the length of word
// Space O(m*n+1)
func exist79(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if helper79(i, j, 0, word, board) {
				return true
			}
		}
	}
	return false
}

func helper79(m, n int, index int, word string, board [][]byte) bool {
	if m == len(board) || m == -1 || n == len(board[0]) || n == -1 {
		return false
	}

	if board[m][n] != word[index] {
		return false
	}

	if index == len(word)-1 {
		return true
	}

	char := board[m][n]
	board[m][n] = 0 // important: temporaily remove duplication
	res := helper79(m+1, n, index+1, word, board) ||
		helper79(m-1, n, index+1, word, board) ||
		helper79(m, n+1, index+1, word, board) ||
		helper79(m, n-1, index+1, word, board)
	board[m][n] = char
	return res
}
