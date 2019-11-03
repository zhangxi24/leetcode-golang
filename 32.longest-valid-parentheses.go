/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 *
 * https://leetcode.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (26.53%)
 * Likes:    2411
 * Dislikes: 110
 * Total Accepted:    224.4K
 * Total Submissions: 844.4K
 * Testcase Example:  '"(()"'
 *
 * Given a string containing just the characters '(' and ')', find the length
 * of the longest valid (well-formed) parentheses substring.
 *
 * Example 1:
 *
 *
 * Input: "(()"
 * Output: 2
 * Explanation: The longest valid parentheses substring is "()"
 *
 *
 * Example 2:
 *
 *
 * Input: ")()())"
 * Output: 4
 * Explanation: The longest valid parentheses substring is "()()"
 *
 *
 */

// @lc code=start
func longestValidParentheses(s string) int {
	q := 0
	l := []int{}
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "(" { // push
			l = append(l, 0)
			if q < 0 {
				q = 1
			} else {
				q++
			}
		} else { // pop
			if q <= 0 {
				l = append(l, 0)
			} else { // 可以配对
				add := l[i-1] + 2
				step := l[i-1] + 2
				for j := i; j > 0 && j-step > 0 && l[j-step] > 0; {
					add += l[j-step]
					j = j - step
					step = l[j]
				}
				l = append(l, add)
			}
			q--
		}
	}

	max := 0
	for i := 0; i < len(s); i++ {
		if l[i] > max {
			max = l[i]
		}
	}
	return max
}

// @lc code=end

