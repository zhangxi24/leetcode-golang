/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 *
 * https://leetcode.com/problems/decode-ways/description/
 *
 * algorithms
 * Medium (23.03%)
 * Likes:    1858
 * Dislikes: 2104
 * Total Accepted:    315.7K
 * Total Submissions: 1.4M
 * Testcase Example:  '"12"'
 *
 * A message containing letters from A-Z is being encoded to numbers using the
 * following mapping:
 *
 *
 * 'A' -> 1
 * 'B' -> 2
 * ...
 * 'Z' -> 26
 *
 *
 * Given a non-empty string containing only digits, determine the total number
 * of ways to decode it.
 *
 * Example 1:
 *
 *
 * Input: "12"
 * Output: 2
 * Explanation: It could be decoded as "AB" (1 2) or "L" (12).
 *
 *
 * Example 2:
 *
 *
 * Input: "226"
 * Output: 3
 * Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2
 * 6).
 *
 */

// @lc code=start
func numDecodings(s string) int {
	if s == "0" {
		return 0
	}
	l := len(s)
	dp := make([]int, l+1)
	dp[l], dp[l-1] = 1, 1
	if s[l-1] == '0' {
		dp[l-1] = 0
	}
	for i := l - 2; i >= 0; i-- {
		if s[i] == '0' {
			continue
		}
		if s[i:i+2] <= "26" {
			dp[i] = dp[i+1] + dp[i+2]
		} else {
			dp[i] = dp[i+1]
		}
	}
	return dp[0]
}

// @lc code=end

