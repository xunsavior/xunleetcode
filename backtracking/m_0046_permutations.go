package backtracking

/*
Company: Microsoft(12), Amazon(8), Facebook(6), Apple(6), Google(5)

Given a collection of distinct integers, return all possible permutations.

Example:
Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

*/

func permute46(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := [][]int{}
	helper46(nums, []int{}, &res)
	return res
}

func helper46(nums, temp []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, temp)
		return
	}

	for i := range nums {
		tmp := append(temp, nums[i])
		tl := len(tmp)
		remains := []int{}
		for j := 0; j < len(nums); j++ {
			if j != i {
				remains = append(remains, nums[j])
			}
		}
		helper46(remains[:len(nums)-1:len(nums)-1], tmp[:tl:tl], res)
	}
}
