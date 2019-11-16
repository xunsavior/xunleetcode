package tree

/*
	Given n, how many structurally unique BST's (binary search trees) that store values 1 ... n?

	Example:
	Input: 3
	Output: 5

	Explanation:
	Given n = 3, there are a total of 5 unique BST's:

	   1         3     3      2      1
		\       /     /      / \      \
		 3     2     1      1   3      2
		/     /       \                 \
	   2     1         2                 3
*/

/*
	f(0) = 1
	f(1) = f(0)*f(0) = 1
	f(2) = f(1)*f(0) + f(0)*f(1) = 1 + 1 = 2
	f(3) = f(2)*f(0) + f(1)*f(1) + f(0)*f(2) = 2 + 1 + 2 = 5
	f(4) = f(3)*f(0) + f(2)*f(1) + f(1)*f(2) + f(0)*f(3) = 5 + 2 + 2 + 5 = 14
	f(5) = f(4)*f(0) + f(3)*f(1) + f(2)*f(2) + f(1)*f(3) + f(0)*f(4) = 14 + 5 + 4 + 5 + 14 = 42
	......
	f(n) = f(n-1)*f(0) + f(n-2)*f(1) + f(n-3)*f(2) + ... + f(2)*f(n-3) + f(1)*f(n-2) + f(0)*f(n-1)
*/

// NumTreesWithDPSolution ...
func NumTreesWithDPSolution(n int) int {
	// create a slice to store each node
	cache := make([]int, n+1)
	// f(0) = 1
	cache[0] = 1
	return getTotalTrees(n, cache)
}

func getTotalTrees(n int, cache []int) int {
	//check if the node has already been stored
	if cache[n] != 0 {
		return cache[n]
	}
	sum := 0
	// implement formular f(n)
	for l, r := n-1, 0; l >= 0 && r <= n-1; l, r = l-1, r+1 {
		sum += getTotalTrees(l, cache) * getTotalTrees(r, cache)
	}
	// need to store every sum
	cache[n] = sum

	return sum
}
