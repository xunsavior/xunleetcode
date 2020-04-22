package greedy

/*
Company: nil; nil; Amazon(2), Druva(1)

Given an array A of integers, we must modify the array in the following way: we choose
an i and replace A[i] with -A[i], and we repeat this process K times in total.  (We may choose the same index i multiple times.)

Return the largest possible sum of the array after modifying it in this way.

Example 1:
Input: A = [4,2,3], K = 1
Output: 5
Explanation: Choose indices (1,) and A becomes [4,-2,3].

Example 2:
Input: A = [3,-1,0,2], K = 3
Output: 6
Explanation: Choose indices (1, 2, 2) and A becomes [3,1,0,2].

Example 3:
Input: A = [2,-3,-1,5,-4], K = 2
Output: 13
Explanation: Choose indices (1, 4) and A becomes [2,3,-1,5,4].
*/

func largestSumAfterKNegations1005(A []int, K int) int {
	negatives, positives := []int{}, []int{}
	for _, v := range A {
		if v < 0 {
			helper1005(&negatives, v)
		} else {
			helper1005(&positives, v)
		}
	}

	res := 0
	for i := 1; i <= K; i++ {
		if len(negatives) > 0 {
			helper1005(&positives, -negatives[0])
			negatives = negatives[1:]
		} else {
			helper1005(&negatives, -positives[0])
			positives = positives[1:]
		}
	}

	for _, v := range negatives {
		res += v
	}

	for _, v := range positives {
		res += v
	}
	return res
}

func helper1005(A *[]int, target int) {
	index, start, end := 0, 0, len(*A)-1

	if len(*A) == 0 || target >= (*A)[len(*A)-1] {
		*A = append(*A, target)
		return
	}

	if target <= (*A)[0] {
		index = start
	} else {
		for start < end {
			mid := start + (end-start)/2
			if (*A)[mid] < target {
				start = mid + 1
			} else {
				end = mid
			}
		}
		index = end
	}

	*A = append(*A, (*A)[len(*A)-1])
	copy((*A)[index+1:], (*A)[index:len((*A))-1])
	(*A)[index] = target
}
