/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 *
 * https://leetcode.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (57.56%)
 * Likes:    3480
 * Dislikes: 206
 * Total Accepted:    411.9K
 * Total Submissions: 713.4K
 * Testcase Example:  '3'
 *
 *
 * Given n pairs of parentheses, write a function to generate all combinations
 * of well-formed parentheses.
 *
 *
 *
 * For example, given n = 3, a solution set is:
 *
 *
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 *
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := []string{}
	helper(n, 0, "", &res)
	return res
}

func helper(n, m int, str string, res *[]string) {
	if n == 0 && m == 0 {
		*res = append(*res, str)
	}
	if n > 0 {
		helper(n-1, m+1, str+"(", res)
	}
	if m > 0 {
		helper(n, m-1, str+")", res)
	}

	return
}

// @lc code=end

