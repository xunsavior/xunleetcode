package stack

/*
Company: Amazon(10), Apple(5), LinkedIn(3), Google(2), Microsoft(2), Uber(2)

Given a list of daily temperatures T, return a list such that, for each
day in the input, tells you how many days you would have to wait until
a warmer temperature. If there is no future day for which this is possible,
put 0 instead.

For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].

Note: The length of temperatures will be in the range [1, 30000]. Each temperature will be an integer in the range [30, 100].
*/

func dailyTemperatures739(T []int) []int {
	stack, res := make([]int, 0), make([]int, len(T))
	for i, v := range T {
		for len(stack) != 0 && v > T[stack[len(stack)-1]] {
			indexDif := i - stack[len(stack)-1]
			res[stack[len(stack)-1]] = indexDif
			stack = stack[:len(stack)-1]
		}
		// store the index instead of val
		stack = append(stack, i)
	}
	return res
}
