import "math"

/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 *
 * https://leetcode.com/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (16.16%)
 * Likes:    815
 * Dislikes: 3935
 * Total Accepted:    222.5K
 * Total Submissions: 1.4M
 * Testcase Example:  '10\n3'
 *
 * Given two integers dividend and divisor, divide two integers without using
 * multiplication, division and mod operator.
 *
 * Return the quotient after dividing dividend by divisor.
 *
 * The integer division should truncate toward zero.
 *
 * Example 1:
 *
 *
 * Input: dividend = 10, divisor = 3
 * Output: 3
 *
 * Example 2:
 *
 *
 * Input: dividend = 7, divisor = -3
 * Output: -2
 *
 * Note:
 *
 *
 * Both dividend and divisor will be 32-bit signed integers.
 * The divisor will never be 0.
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose
 * of this problem, assume that your function returns 2^31 − 1 when the
 * division result overflows.
 *
 *
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	res := 0
	a := dividend
	if dividend < 0 {
		a = -dividend
	}
	b := divisor
	if divisor < 0 {
		b = -divisor
	}
	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}

	for a >= b {
		a -= b
		res++
	}

	if sign == -1 {
		res = -res
	}
	if res < math.MinInt32 || res > math.MaxInt32 {
		res = math.MaxInt32
	}

	return res
}

// @lc code=end

