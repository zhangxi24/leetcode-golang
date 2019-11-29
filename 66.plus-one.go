/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 *
 * https://leetcode.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (41.73%)
 * Likes:    1052
 * Dislikes: 1831
 * Total Accepted:    466.4K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,3]'
 *
 * Given a non-empty array of digits representing a non-negative integer, plus
 * one to the integer.
 *
 * The digits are stored such that the most significant digit is at the head of
 * the list, and each element in the array contain a single digit.
 *
 * You may assume the integer does not contain any leading zero, except the
 * number 0 itself.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3]
 * Output: [1,2,4]
 * Explanation: The array represents the integer 123.
 *
 *
 * Example 2:
 *
 *
 * Input: [4,3,2,1]
 * Output: [4,3,2,2]
 * Explanation: The array represents the integer 4321.
 *
 */

// @lc code=start
func plusOne(digits []int) []int {
	up, s := 1, 0
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		s = up + digits[i]
		up = s / 10
		digits[i] = s % 10
	}
	if up == 0 {
		return digits
	} else {
		res := []int{up}
		res = append(res, digits...)
		return res
	}
}

// @lc code=end

