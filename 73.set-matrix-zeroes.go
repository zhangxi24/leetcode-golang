/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 *
 * https://leetcode.com/problems/set-matrix-zeroes/description/
 *
 * algorithms
 * Medium (41.10%)
 * Likes:    1381
 * Dislikes: 237
 * Total Accepted:    250.2K
 * Total Submissions: 606K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * Given a m x n matrix, if an element is 0, set its entire row and column to
 * 0. Do it in-place.
 *
 * Example 1:
 *
 *
 * Input:
 * [
 * [1,1,1],
 * [1,0,1],
 * [1,1,1]
 * ]
 * Output:
 * [
 * [1,0,1],
 * [0,0,0],
 * [1,0,1]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input:
 * [
 * [0,1,2,0],
 * [3,4,5,2],
 * [1,3,1,5]
 * ]
 * Output:
 * [
 * [0,0,0,0],
 * [0,4,5,0],
 * [0,3,1,0]
 * ]
 *
 *
 * Follow up:
 *
 *
 * A straight forward solution using O(mn) space is probably a bad idea.
 * A simple improvement uses O(m + n) space, but still not the best
 * solution.
 * Could you devise a constant space solution?
 *
 *
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	m := len(matrix)
	n := len(matrix[0])
	zerocol := map[int]bool{}
	zerorow := map[int]bool{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zerocol[i] = true
				zerorow[j] = true
			}
		}
	}
	for k := range zerocol {
		for i := 0; i < n; i++ {
			matrix[k][i] = 0
		}
	}
	for k := range zerorow {
		for i := 0; i < m; i++ {
			matrix[i][k] = 0
		}
	}
}

// @lc code=end

