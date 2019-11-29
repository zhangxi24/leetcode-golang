/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 *
 * https://leetcode.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (49.16%)
 * Likes:    636
 * Dislikes: 90
 * Total Accepted:    159.2K
 * Total Submissions: 321K
 * Testcase Example:  '3'
 *
 * Given a positive integer n, generate a square matrix filled with elements
 * from 1 to n^2 in spiral order.
 *
 * Example:
 *
 *
 * Input: 3
 * Output:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 8, 9, 4 ],
 * ⁠[ 7, 6, 5 ]
 * ]
 *
 *
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	res := [][]int{}
	if n == 0 {
		return res
	}
	for i := 0; i < n; i++ {
		tmp := make([]int, n)
		res = append(res, tmp)
	}
	round := (n + 1) / 2
	num := 1
	for i := 0; i < round; i++ {
		for a := i; a < n-i-1; a++ {
			res[i][a] = num
			num++
		}
		for a := i; a < n-i-1; a++ {
			res[a][n-i-1] = num
			num++
		}
		for a := n - i - 1; a > i; a-- {
			res[n-i-1][a] = num
			num++
		}
		for a := n - i - 1; a > i; a-- {
			res[a][i] = num
			num++
		}
		if i == n-i-1 {
			res[i][i] = num
		}
	}

	return res
}

// @lc code=end

