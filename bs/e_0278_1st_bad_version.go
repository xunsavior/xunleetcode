package bs

/*
Company: Facebook(41); Adobe(3), Google(2), Microsoft(2); Bloomberg(2), Uber(2)

You are a product manager and currently leading a team to develop a new product.
Unfortunately, the latest version of your product fails the quality check. Since
each version is developed based on the previous version, all the versions after
a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API bool isBadVersion(version) which will return whether version is bad. Implement a function to find the first bad version.
You should minimize the number of calls to the API.

Example:
Given n = 5, and version = 4 is the first bad version.
call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true

Then 4 is the first bad version.
*/

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion278(n int) int {
	start, end := 1, n

	for start < end {
		mid := start + (end-start)/2

		if !isBadVersion(mid) {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return start
}

func isBadVersion(n int) bool {
	return false
}
