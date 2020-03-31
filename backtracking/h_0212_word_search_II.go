package backtracking

/*
Company: Amazon(20), Microsoft(7), Google(5), Facebook(3), Snapchat(3)

Given a 2D board and a list of words from the dictionary, find all words in the board.

Each word must be constructed from letters of sequentially adjacent cell, where "adjacent"
cells are those horizontally or vertically neighboring. The same letter cell may not be
used more than once in a word.

Example:
Input:
board = [
  ['o','a','a','n'],
  ['e','t','a','e'],
  ['i','h','k','r'],
  ['i','f','l','v']
]
words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]

Note:
All inputs are consist of lowercase letters a-z.
The values of words are distinct.
*/

func findWords212(board [][]byte, words []string) []string {
	if len(words) == 0 || len(board) == 0 {
		return nil
	}

	dict, res := make(map[string]bool), []string{}

	for _, word := range words {
		for i := range board {
			for j := range board[i] {
				if helper212(i, j, 0, board, word) && dict[word] == false {
					res = append(res, word)
					dict[word] = true
				}
			}
		}
	}

	return res
}

// DFS
func helper212(m, n, index int, board [][]byte, word string) bool {
	if m == -1 || m == len(board) || n == -1 || n == len(board[0]) || word[index] != board[m][n] {
		return false
	}

	if len(word)-1 == index {
		return true
	}

	val := board[m][n]
	board[m][n] = 48
	res := helper212(m+1, n, index+1, board, word) ||
		helper212(m-1, n, index+1, board, word) ||
		helper212(m, n+1, index+1, board, word) ||
		helper212(m, n-1, index+1, board, word)
	board[m][n] = val
	return res
}
