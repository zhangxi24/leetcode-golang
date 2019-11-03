/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (28.67%)
 * Likes:    1014
 * Dislikes: 2441
 * Total Accepted:    374.4K
 * Total Submissions: 1.3M
 * Testcase Example:  '2.00000\n10'
 *
 * Implement pow(x, n), which calculates x raised to the power n (x^n).
 *
 * Example 1:
 *
 *
 * Input: 2.00000, 10
 * Output: 1024.00000
 *
 *
 * Example 2:
 *
 *
 * Input: 2.10000, 3
 * Output: 9.26100
 *
 *
 * Example 3:
 *
 *
 * Input: 2.00000, -2
 * Output: 0.25000
 * Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
 *
 *
 * Note:
 *
 *
 * -100.0 < x < 100.0
 * n is a 32-bit signed integer, within the range [−2^31, 2^31 − 1]
 *
 *
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	res := 1.0
	if n == 0 {
		return res
	}

	negetive := false
	if n < 0 {
		negetive = true
		n = -n
	}

	res = pow(x, n)
	if negetive == true {
		res = 1 / res
	}

	return res
}

func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	}

	if n%2 == 0 {
		tmp := pow(x, n/2)
		return tmp * tmp
	} else {
		tmp := pow(x, n/2)
		return tmp * tmp * x
	}
}

// @lc code=end

