/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (45.26%)
 * Likes:    1683
 * Dislikes: 1404
 * Total Accepted:    710.9K
 * Total Submissions: 1.6M
 * Testcase Example:  '121'
 *
 * Determine whether an integer is a palindrome. An integer is a palindrome
 * when it reads the same backward as forward.
 *
 * Example 1:
 *
 *
 * Input: 121
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 *
 *
 * Follow up:
 *
 * Coud you solve it without converting the integer to a string?
 *
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	b := 0
	for a := x; a > 0; a /= 10 {
		b = b*10 + a%10
	}
	if x == b {
		return true
	} else {
		return false
	}
}

// @lc code=end

