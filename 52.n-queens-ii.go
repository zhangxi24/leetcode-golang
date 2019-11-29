/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 *
 * https://leetcode.com/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (54.12%)
 * Likes:    341
 * Dislikes: 129
 * Total Accepted:    112.7K
 * Total Submissions: 207.5K
 * Testcase Example:  '4'
 *
 * The n-queens puzzle is the problem of placing n queens on an n×n chessboard
 * such that no two queens attack each other.
 *
 *
 *
 * Given an integer n, return the number of distinct solutions to the n-queens
 * puzzle.
 *
 * Example:
 *
 *
 * Input: 4
 * Output: 2
 * Explanation: There are two distinct solutions to the 4-queens puzzle as
 * shown below.
 * [
 * [".Q..",  // Solution 1
 * "...Q",
 * "Q...",
 * "..Q."],
 *
 * ["..Q.",  // Solution 2
 * "Q...",
 * "...Q",
 * ".Q.."]
 * ]
 *
 *
 */

// @lc code=start
func totalNQueens(n int) int {
	num := 0
	used := [][]bool{}
	for i := 0; i < n; i++ {
		tmp := make([]bool, n)
		used = append(used, tmp)
	}
	helper(&num, n, 0, used, []string{})
	return num
}

func helper(num *int, depth, i int, used [][]bool, r []string) {
	if i == depth {
		*num++
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
				helper(num, depth, i+1, used, r)
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

