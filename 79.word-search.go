/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.com/problems/word-search/description/
 *
 * algorithms
 * Medium (32.66%)
 * Likes:    2427
 * Dislikes: 130
 * Total Accepted:    354.1K
 * Total Submissions: 1.1M
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * Given a 2D board and a word, find if the word exists in the grid.
 *
 * The word can be constructed from letters of sequentially adjacent cell,
 * where "adjacent" cells are those horizontally or vertically neighboring. The
 * same letter cell may not be used more than once.
 *
 * Example:
 *
 *
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 *
 * Given word = "ABCCED", return true.
 * Given word = "SEE", return true.
 * Given word = "ABCB", return false.
 *
 *
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 {
		return false
	}

	pos := [][]int{}
	m, n := len(board), len(board[0])
	used := [][]bool{}
	for i := 0; i < m; i++ {
		tmp := make([]bool, n)
		used = append(used, tmp)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				pos = append(pos, []int{i, j})
			}
		}
	}
	for _, p := range pos { // 尝试每个起始位置
		used[p[0]][p[1]] = true
		if helper(used, 1, word, p[0], p[1], m, n, board) == true {
			return true
		}
		used[p[0]][p[1]] = false
	}
	return false

}

func helper(used [][]bool, k int, word string, i, j, m, n int, board [][]byte) bool {
	if k == len(word) {
		return true
	}
	poss := aroundPos(i, j, m, n)
	for _, p := range poss {
		if used[p[0]][p[1]] == false && board[p[0]][p[1]] == word[k] {
			used[p[0]][p[1]] = true
			if helper(used, k+1, word, p[0], p[1], m, n, board) == true {
				return true
			}
			used[p[0]][p[1]] = false
		}
	}
	return false
}

func aroundPos(i, j, m, n int) [][]int {
	res := [][]int{}
	if i != 0 { // up
		res = append(res, []int{i - 1, j})
	}
	if j != 0 { // right
		res = append(res, []int{i, j - 1})
	}
	if i != m-1 { // down
		res = append(res, []int{i + 1, j})
	}
	if j != n-1 { // left
		res = append(res, []int{i, j + 1})
	}
	return res
}

// @lc code=end

