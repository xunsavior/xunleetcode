package backtracking

/*
Company: Facebook(2), Apple(2), Oracle(2), Amazon(2); Microsoft(4), Google(2)

The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.
Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.

Example:
Input: 4
Output: [
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.
*/

func solveNQueens51(n int) [][]string {
	res := [][]string{}
	if n == 0 {
		return res
	}
	helper51(0, n, []int{}, &res)
	return res
}

func helper51(index, n int, queens []int, res *[][]string) {
	if index == n {
		strs := []string{}
		for _, v := range queens {
			row := ""
			for j := 0; j < n; j++ {
				row += "."
			}
			bytes := []byte(row)
			bytes[v] = 81
			row = string(bytes)
			strs = append(strs, row)
		}
		*res = append(*res, strs)
		return
	}

	dict := make(map[int]bool)
	for i, v := range queens {
		dict[v] = true
		dict[v+len(queens)-i] = true
		dict[v-len(queens)+i] = true
	}

	for i := 0; i < n; i++ {
		if dict[i] == true {
			continue
		}
		newQueue := append(queens, i)
		helper51(index+1, n, newQueue[:len(newQueue):len(newQueue)], res)
	}
}
