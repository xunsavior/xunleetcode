package tree

import "math"

/*
Company: Google(5), Amazon(3), Uber(2); nil; Facebook(2)

The thief has found himself a new place for his thievery again. There is only one entrance to this area,
called the "root." Besides the root, each house has one and only one parent house. After a tour,
the smart thief realized that "all houses in this place forms a binary tree". It will automatically contact
the police if two directly-linked houses were broken into on the same night.

Determine the maximum amount of money the thief can rob tonight without alerting the police.

Example 1:
Input: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \
     3   1

Output: 7
Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.

Example 2:
Input: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \
 1   3   1
Output: 9
Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.
*/

type cache struct {
	node   *TNode
	isDone bool
}

func rob337(root *TNode) int {
	dp := make(map[cache]int)
	return helper337(false, root, dp)
}

func helper337(isDone bool, root *TNode, dp map[cache]int) int {
	if root == nil {
		return 0
	}

	ca := cache{
		node:   root,
		isDone: isDone,
	}

	if res, ok := dp[ca]; ok {
		return res
	}

	opt1 := helper337(false, root.Left, dp) + helper337(false, root.Right, dp)
	if isDone {
		dp[ca] = opt1
		return opt1
	}

	opt2 := root.Val + helper337(true, root.Left, dp) + helper337(true, root.Right, dp)
	res := int(math.Max(float64(opt1), float64(opt2)))
	dp[ca] = res
	return res
}
