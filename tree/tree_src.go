package tree

// TNode ...
type TNode struct {
	Val   int
	Left  *TNode
	Right *TNode
}

// NodeSum ...
func NodeSum(nodes []*TNode) int {
	sum := 0
	for _, v := range nodes {
		sum += v.Val
	}
	return sum
}
