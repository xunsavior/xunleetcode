package dfsandbfs

/*
Company: Two Sigma(18), Pocket Gems(13), Twitter(7), Amazon(6), Google(2)

There are N students in a class. Some of them are friends, while some are not.
Their friendship is transitive in nature. For example, if A is a direct friend
of B, and B is a direct friend of C, then A is an indirect friend of C.

We defined a friend circle is a group of students who are direct or indirect friends.

Given a N*N matrix M representing the friend relationship between students in the class.
If M[i][j] = 1, then the ith and jth students are direct friends with each other, otherwise not.
And you have to output the total number of friend circles among all the students.

Example 1:
Input:
[[1,1,0],
 [1,1,0],
 [0,0,1]]
Output: 2
Explanation:The 0th and 1st students are direct friends, so they are in a friend circle.
The 2nd student himself is in a friend circle. So return 2.

Example 2:
Input:
[[1,1,0],
 [1,1,1],
 [0,1,1]]
Output: 1
Explanation:The 0th and 1st students are direct friends, the 1st and 2nd students are direct friends,
so the 0th and 2nd students are indirect friends. All of them are in the same friend circle, so return 1.

Note:
N is in range [1,200].
M[i][i] = 1 for all students.
If M[i][j] = 1, then M[j][i] = 1.
*/

func findCircleNum547(M [][]int) int {
	// the total number of students is the length of M
	// we create a slice to store person have been visited
	visitedPerson := make([]int, len(M))
	res := 0
	// we find all the direct and indirect friends of person i that has NOT been visited
	for i := range M {
		if visitedPerson[i] == 0 {
			helper547(i, visitedPerson, M)
			res++
		}
	}
	return res
}

// DFS
func helper547(i int, visited []int, M [][]int) {
	for j := range M {
		if M[i][j] == 1 && visited[j] == 0 {
			visited[j] = 1
			helper547(j, visited, M)
		}
	}
}
