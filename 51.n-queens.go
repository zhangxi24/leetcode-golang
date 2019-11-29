/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 *
 * https://leetcode.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (41.87%)
 * Likes:    1239
 * Dislikes: 57
 * Total Accepted:    165K
 * Total Submissions: 391.5K
 * Testcase Example:  '4'
 *
 * The n-queens puzzle is the problem of placing n queens on an n×n chessboard
 * such that no two queens attack each other.
 *
 *
 *
 * Given an integer n, return all distinct solutions to the n-queens puzzle.
 *
 * Each solution contains a distinct board configuration of the n-queens'
 * placement, where 'Q' and '.' both indicate a queen and an empty space
 * respectively.
 *
 * Example:
 *
 *
 * Input: 4
 * Output: [
 * ⁠[".Q..",  // Solution 1
 * ⁠ "...Q",
 * ⁠ "Q...",
 * ⁠ "..Q."],
 *
 * ⁠["..Q.",  // Solution 2
 * ⁠ "Q...",
 * ⁠ "...Q",
 * ⁠ ".Q.."]
 * ]
 * Explanation: There exist two distinct solutions to the 4-queens puzzle as
 * shown above.
 *
 *
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	res := [][]string{}
	used := [][]bool{}
	for i := 0; i < n; i++ {
		tmp := make([]bool, n)
		used = append(used, tmp)
	}
	helper(&res, n, 0, used, []string{})
	return res
}

func helper(res *[][]string, depth, i int, used [][]bool, r []string) {
	if i == depth {
		tmp := make([]string, len(r))
		copy(tmp, r)
		*res = append(*res, tmp)
		return
	}

	for j := 0; j < depth; j++ {
		if used[i][j] == false {
			used[i][j] = true
			if isvalid(r, i, j, depth) {
				add := ""
				for n := 0; n < depth; n++ {
					if n == j {
						add += "Q"
					} else {
						add += "."
					}
				}
				r = append(r, add)
				helper(res, depth, i+1, used, r)
				r = r[:len(r)-1]
			}
			used[i][j] = false
		}
	}

}

func isvalid(r []string, level, index, depth int) bool {
	if len(r) == 0 {
		return true
	}

	// 同一列不能放置
	for i := 0; i < len(r); i++ {
		if r[i][index] == 'Q' {
			return false
		}
	}
	// 左上不能放置
	for m, n := level-1, index-1; m >= 0 && n >= 0; {
		if r[m][n] == 'Q' {
			return false
		}
		m--
		n--
	}

	// 右上不能放置
	for m, n := level-1, index+1; m >= 0 && n < depth; {
		if r[m][n] == 'Q' {
			return false
		}
		m--
		n++
	}

	return true
}

// @lc code=end

