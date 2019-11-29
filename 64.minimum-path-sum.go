/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 *
 * https://leetcode.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (49.04%)
 * Likes:    1767
 * Dislikes: 45
 * Total Accepted:    277.5K
 * Total Submissions: 561.5K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * Given a m x n grid filled with non-negative numbers, find a path from top
 * left to bottom right which minimizes the sum of all numbers along its path.
 *
 * Note: You can only move either down or right at any point in time.
 *
 * Example:
 *
 *
 * Input:
 * [
 *  [1,3,1],
 * ⁠ [1,5,1],
 * ⁠ [4,2,1]
 * ]
 * Output: 7
 * Explanation: Because the path 1→3→1→1→1 minimizes the sum.
 *
 *
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])
	dp := [][]int{}
	for i := 0; i < m; i++ {
		tmp := make([]int, n)
		dp = append(dp, tmp)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			min := dp[i-1][j]
			if dp[i][j-1] < dp[i-1][j] {
				min = dp[i][j-1]
			}
			dp[i][j] = min + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

// @lc code=end

