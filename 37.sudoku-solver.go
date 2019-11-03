/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 *
 * https://leetcode.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (39.18%)
 * Likes:    1165
 * Dislikes: 76
 * Total Accepted:    149.8K
 * Total Submissions: 380.9K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * Write a program to solve a Sudoku puzzle by filling the empty cells.
 *
 * A sudoku solution must satisfy all of the following rules:
 *
 *
 * Each of the digits 1-9 must occur exactly once in each row.
 * Each of the digits 1-9 must occur exactly once in each column.
 * Each of the the digits 1-9 must occur exactly once in each of the 9 3x3
 * sub-boxes of the grid.
 *
 *
 * Empty cells are indicated by the character '.'.
 *
 *
 * A sudoku puzzle...
 *
 *
 * ...and its solution numbers marked in red.
 *
 * Note:
 *
 *
 * The given board contain only digits 1-9 and the character '.'.
 * You may assume that the given Sudoku puzzle will have a single unique
 * solution.
 * The given board size is always 9x9.
 *
 *
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	nums := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	helper(&board, 0, 0, &nums)
}

func helper(board *[][]byte, r, c int, nums *[]byte) bool {
	if c == 9 {
		c = 0
		r += 1
	}
	if r == 9 {
		return true
	}
	if (*board)[r][c] != '.' {
		return helper(board, r, c+1, nums)
	}
	for _, n := range *nums {
		if isValid(board, r, c, n) {
			(*board)[r][c] = n
			if helper(board, r, c+1, nums) {
				return true
			}
		}
	}
	(*board)[r][c] = '.'
	return false
}

func isValid(board *[][]byte, r, c int, value byte) bool {
	for i := 0; i < 9; i++ {
		if (*board)[i][c] == value {
			return false
		}
		if (*board)[r][i] == value {
			return false
		}
	}

	r0, c0 := r/3, c/3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (*board)[i+r0*3][j+c0*3] == value {
				return false
			}
		}
	}
	return true
}

// @lc code=end

