package hashtable

/*
Company: Microsoft(4), Amazon(5), Apple(2), Google(2)
Count the number of prime numbers less than a non-negative number, n.

Example:
Input: 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
*/

func countPrimes204(n int) int {
	res := 0
	nonPrimeDict := make(map[int]bool)
	for i := 2; i < n; i++ {
		if !nonPrimeDict[i] {
			res++
			for j := 2; i*j < n; j++ {
				nonPrimeDict[i*j] = true
			}
		}
	}
	return res
}
