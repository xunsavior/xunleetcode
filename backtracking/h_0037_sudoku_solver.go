package backtracking

/*
Company: Microsoft(5), Amazon(4), Facebook(2), Oracle(2); null; Google(6)

Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

Each of the digits 1-9 must occur exactly once in each row.
Each of the digits 1-9 must occur exactly once in each column.
Each of the the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
Empty cells are indicated by the character '.'.


A sudoku puzzle...


...and its solution numbers marked in red.

Note:
The given board contain only digits 1-9 and the character '.'.
You may assume that the given Sudoku puzzle will have a single unique solution.
The given board size is always 9x9.
*/

func solveSudoku37(board [][]byte) {
	helper37(board)
}

func helper37(board [][]byte) bool {
	for i := range board {
		for j := range board[0] {
			if string(board[i][j]) == "." {

				for code := 49; code < 58; code++ {
					if isValid37(i, j, code, board) { // we need to check if it is valid or not BEFORE updating the value of slot
						board[i][j] = byte(code)
						if helper37(board) { // we need to check if it is doable after updating the value of this slot
							return true
						}
						board[i][j] = 46
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid37(m, n, val int, board [][]byte) bool {
	for _, v := range board[m] {
		if int(v) == val {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		v := int(board[i][n])
		if v == val {
			return false
		}
	}

	x, y := m/3, n/3

	for i := 3 * x; i < 3*(x+1); i++ {
		for j := 3 * y; j < 3*(y+1); j++ {
			if val == int(board[i][j]) {
				return false
			}
		}
	}

	return true
}
