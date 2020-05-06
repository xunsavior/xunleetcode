package design

/*
Company: Amazon(11), Microsoft(9), TripleByte(8), Facebook(7), Apple(3)

Design a Tic-tac-toe game that is played between two players on a n x n grid.

You may assume the following rules:

A move is guaranteed to be valid and is placed on an empty block.
Once a winning condition is reached, no more moves is allowed.
A player who succeeds in placing n of their marks in a horizontal, vertical, or diagonal row wins the game.

Follow up:
Could you do better than O(n2) per move() operation?
*/

type TicTacToe struct {
	Board [][]int
}

/** Initialize your data structure here. */
func Constructor348(n int) TicTacToe {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	return TicTacToe{
		Board: res,
	}
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
func (this *TicTacToe) Move348(row int, col int, player int) int {
	if isWinning(this.Board, row, col, player) {
		return player
	}
	return 0
}

func isWinning(board [][]int, row, col, player int) bool {
	board[row][col] = player
	for i := 0; i < len(board); i++ {
		if board[i][col] != player {
			break
		}
		if i == len(board)-1 {
			return true
		}
	}

	for i := 0; i < len(board); i++ {
		if board[row][i] != player {
			break
		}
		if i == len(board)-1 {
			return true
		}
	}

	if row == col {
		for i, j := 0, 0; i < len(board) && j < len(board); i, j = i+1, j+1 {
			if board[i][j] != player {
				break
			}
			if i == len(board)-1 {
				return true
			}
		}
	}

	if row+col == len(board)-1 {
		for i, j := 0, len(board)-1; i < len(board) && j >= 0; i, j = i+1, j-1 {
			if board[i][j] != player {
				break
			}
			if i == len(board)-1 {
				return true
			}
		}
	}

	return false
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
