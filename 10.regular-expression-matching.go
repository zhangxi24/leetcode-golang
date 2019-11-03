/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 *
 * https://leetcode.com/problems/regular-expression-matching/description/
 *
 * algorithms
 * Hard (25.82%)
 * Likes:    3204
 * Dislikes: 599
 * Total Accepted:    353.8K
 * Total Submissions: 1.4M
 * Testcase Example:  '"aa"\n"a"'
 *
 * Given an input string (s) and a pattern (p), implement regular expression
 * matching with support for '.' and '*'.
 *
 *
 * '.' Matches any single character.
 * '*' Matches zero or more of the preceding element.
 *
 *
 * The matching should cover the entire input string (not partial).
 *
 * Note:
 *
 *
 * s could be empty and contains only lowercase letters a-z.
 * p could be empty and contains only lowercase letters a-z, and characters
 * like . or *.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * s = "aa"
 * p = "a"
 * Output: false
 * Explanation: "a" does not match the entire string "aa".
 *
 *
 * Example 2:
 *
 *
 * Input:
 * s = "aa"
 * p = "a*"
 * Output: true
 * Explanation: '*' means zero or more of the preceding element, 'a'.
 * Therefore, by repeating 'a' once, it becomes "aa".
 *
 *
 * Example 3:
 *
 *
 * Input:
 * s = "ab"
 * p = ".*"
 * Output: true
 * Explanation: ".*" means "zero or more (*) of any character (.)".
 *
 *
 * Example 4:
 *
 *
 * Input:
 * s = "aab"
 * p = "c*a*b"
 * Output: true
 * Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore,
 * it matches "aab".
 *
 *
 * Example 5:
 *
 *
 * Input:
 * s = "mississippi"
 * p = "mis*is*p*."
 * Output: false
 *
 *
 */

// @lc code=start
func isMatch(s string, p string) bool {
	slen, plen := len(s), len(p)
	var dp [][]bool
	for i := 0; i <= slen; i++ {
		t := make([]bool, plen+1)
		dp = append(dp, t)
	}

	// 0 代表空串
	for i := 0; i <= slen; i++ {
		for j := 0; j <= plen; j++ {
			// init
			if i == 0 && j == 0 {
				dp[i][j] = true
				continue
			} else if i == 0 {
				if (j-1)%2 == 1 && p[j-1] == '*' {
					dp[i][j] = dp[i][j-2]
				}
				continue
			} else if j == 0 {
				dp[i][j] = false
				continue
			}

			// 转移方程
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					dp[i][j] = dp[i-1][j]
				}
				if dp[i][j-2] == true {
					dp[i][j] = true
				}
			} else {
				dp[i][j] = false
			}

		}
	}

	return dp[slen][plen]
}

// @lc code=end

