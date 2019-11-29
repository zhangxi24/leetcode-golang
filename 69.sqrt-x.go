/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 *
 * https://leetcode.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (32.27%)
 * Likes:    940
 * Dislikes: 1576
 * Total Accepted:    438.6K
 * Total Submissions: 1.4M
 * Testcase Example:  '4'
 *
 * Implement int sqrt(int x).
 *
 * Compute and return the square root of x, where x is guaranteed to be a
 * non-negative integer.
 *
 * Since the return type is an integer, the decimal digits are truncated and
 * only the integer part of the result is returned.
 *
 * Example 1:
 *
 *
 * Input: 4
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since
 * the decimal part is truncated, 2 is returned.
 *
 *
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left, right := 1, x/2
	for left <= right {
		mid := (left + right) / 2
		mul := mid * mid
		if mul == x {
			return mid
		} else if mul < x {
			left = mid + 1
		} else if mul > x {
			right = mid - 1
		}
	}
	return left - 1
}

// @lc code=end

