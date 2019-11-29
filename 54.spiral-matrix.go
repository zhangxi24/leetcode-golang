/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 *
 * https://leetcode.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (31.74%)
 * Likes:    1481
 * Dislikes: 458
 * Total Accepted:    282.7K
 * Total Submissions: 886.4K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given a matrix of m x n elements (m rows, n columns), return all elements of
 * the matrix in spiral order.
 *
 * Example 1:
 *
 *
 * Input:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 * Output: [1,2,3,6,9,8,7,4,5]
 *
 *
 * Example 2:
 *
 * Input:
 * [
 * ⁠ [1, 2, 3, 4],
 * ⁠ [5, 6, 7, 8],
 * ⁠ [9,10,11,12]
 * ]
 * Output: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	m := len(matrix)
	n := len(matrix[0])
	res := []int{}
	min := m
	if n < min {
		min = n
	}
	round := (min + 1) / 2
	for i := 0; i < round; i++ { // 层数 (i,i) (i, n-1-i) (m-1-i ,n-1-i) (m-1-i,i)
		for j := i; j <= n-1-i; j++ {
			res = append(res, matrix[i][j])
		}
		for j := i + 1; j <= m-1-i; j++ {
			res = append(res, matrix[j][n-1-i])
		}
		if (m-1-i) > i && (n-1-i) > i {
			for j := n - 2 - i; j >= i; j-- {
				res = append(res, matrix[m-1-i][j])
			}
			for j := m - 2 - i; j > i; j-- {
				res = append(res, matrix[j][i])
			}

		}
	}
	return res
}

// @lc code=end

