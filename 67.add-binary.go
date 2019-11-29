/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 *
 * https://leetcode.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (40.92%)
 * Likes:    1227
 * Dislikes: 229
 * Total Accepted:    356.2K
 * Total Submissions: 864.9K
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings, return their sum (also a binary string).
 *
 * The input strings are both non-empty and contains only characters 1 or 0.
 *
 * Example 1:
 *
 *
 * Input: a = "11", b = "1"
 * Output: "100"
 *
 * Example 2:
 *
 *
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 *
 */

// @lc code=start
func addBinary(a string, b string) string {
	res := ""
	up := 0
	i, j := len(a)-1, len(b)-1
	for i >= 0 && j >= 0 {
		s := int(a[i]-'0') + int(b[j]-'0') + up
		up = s / 2
		res = string('0'+s%2) + res
		i--
		j--
	}
	for i >= 0 {
		s := int(a[i]-'0') + up
		up = s / 2
		res = string('0'+s%2) + res
		i--
	}
	for j >= 0 {
		s := int(b[j]-'0') + up
		up = s / 2
		res = string('0'+s%2) + res
		j--
	}

	if up == 1 {
		res = "1" + res
	}
	return res
}

// @lc code=end

