/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (37.40%)
 * Likes:    3511
 * Dislikes: 171
 * Total Accepted:    735.1K
 * Total Submissions: 2M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 *
 *
 * Note that an empty string is also considered valid.
 *
 * Example 1:
 *
 *
 * Input: "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: "(]"
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: "([)]"
 * Output: false
 *
 *
 * Example 5:
 *
 *
 * Input: "{[]}"
 * Output: true
 *
 *
 */

// @lc code=start
func isValid(s string) bool {
	pool := ""
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "(" || s[i:i+1] == "[" || s[i:i+1] == "{" {
			pool += s[i : i+1]
		} else {
			l := len(pool)
			if l == 0 {
				return false
			}
			if (s[i:i+1] == ")" && pool[l-1:l] == "(") ||
				(s[i:i+1] == "]" && pool[l-1:l] == "[") ||
				(s[i:i+1] == "}" && pool[l-1:l] == "{") {
				pool = pool[:l-1]
			} else {
				return false
			}
		}
	}
	if len(pool) == 0 {
		return true
	} else {
		return false
	}
}

// @lc code=end

