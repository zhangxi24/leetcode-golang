import "math"

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (25.51%)
 * Likes:    2534
 * Dislikes: 3950
 * Total Accepted:    838.6K
 * Total Submissions: 3.3M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 *
 * Example 1:
 *
 *
 * Input: 123
 * Output: 321
 *
 *
 * Example 2:
 *
 *
 * Input: -123
 * Output: -321
 *
 *
 * Example 3:
 *
 *
 * Input: 120
 * Output: 21
 *
 *
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 0 when the reversed
 * integer overflows.
 *
 */

// @lc code=start
func reverse(x int) int {
	ret := 0
	for ; x != 0; x /= 10 {
		ret = ret*10 + x%10
		if ret < -math.MaxInt32-1 || ret > math.MaxInt32 {
			ret = 0
			return ret
		}
	}
	return ret
}

// @lc code=end

