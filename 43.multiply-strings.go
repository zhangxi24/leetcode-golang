/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 *
 * https://leetcode.com/problems/multiply-strings/description/
 *
 * algorithms
 * Medium (31.79%)
 * Likes:    1265
 * Dislikes: 581
 * Total Accepted:    235.3K
 * Total Submissions: 737.7K
 * Testcase Example:  '"2"\n"3"'
 *
 * Given two non-negative integers num1 and num2 represented as strings, return
 * the product of num1 and num2, also represented as a string.
 *
 * Example 1:
 *
 *
 * Input: num1 = "2", num2 = "3"
 * Output: "6"
 *
 * Example 2:
 *
 *
 * Input: num1 = "123", num2 = "456"
 * Output: "56088"
 *
 *
 * Note:
 *
 *
 * The length of both num1 and num2 is < 110.
 * Both num1 and num2 contain only digits 0-9.
 * Both num1 and num2 do not contain any leading zero, except the number 0
 * itself.
 * You must not use any built-in BigInteger library or convert the inputs to
 * integer directly.
 *
 *
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {

			b := int(num1[i]-'0') * int(num2[j]-'0')
			up := 0
			pos := len(num1) - 1 - i + len(num2) - 1 - j
			for b > 0 || up != 0 {
				u := up
				up = (res[pos] + b%10 + up) / 10
				res[pos] = (res[pos] + b%10 + u) % 10
				pos++
				b /= 10
			}
		}
	}

	r := ""
	for i := 0; i < len(res)-1; i++ {
		r = string('0'+res[i]) + r
	}
	if res[len(res)-1] != 0 {
		r = string('0'+res[len(res)-1]) + r
	}

	return r
}

// @lc code=end

