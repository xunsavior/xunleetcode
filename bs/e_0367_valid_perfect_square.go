package bs

/*
Company: nil; LinkedIn(4), Microsoft(2); Google(2), Amazon(2)

Given a positive integer num, write a function which returns True if num is a perfect square else False.

Note: Do not use any built-in library function such as sqrt.

Example 1:
Input: 16
Output: true

Example 2:
Input: 14
Output: false
*/

func isPerfectSquare367(num int) bool {
	if num < 2 {
		return true
	}

	start, end := 2, num

	for start <= end {
		mid := start + (end-start)/2
		sqr := mid * mid
		if num > sqr {
			start = mid + 1
		} else if num < sqr {
			end = mid - 1
		} else {
			return true
		}
	}

	return false
}
