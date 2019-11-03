/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 *
 * https://leetcode.com/problems/wildcard-matching/description/
 *
 * algorithms
 * Hard (23.51%)
 * Likes:    1385
 * Dislikes: 82
 * Total Accepted:    203.2K
 * Total Submissions: 861.3K
 * Testcase Example:  '"aa"\n"a"'
 *
 * Given an input string (s) and a pattern (p), implement wildcard pattern
 * matching with support for '?' and '*'.
 *
 *
 * '?' Matches any single character.
 * '*' Matches any sequence of characters (including the empty sequence).
 *
 *
 * The matching should cover the entire input string (not partial).
 *
 * Note:
 *
 *
 * s could be empty and contains only lowercase letters a-z.
 * p could be empty and contains only lowercase letters a-z, and characters
 * like ? or *.
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
 * p = "*"
 * Output: true
 * Explanation: '*' matches any sequence.
 *
 *
 * Example 3:
 *
 *
 * Input:
 * s = "cb"
 * p = "?a"
 * Output: false
 * Explanation: '?' matches 'c', but the second letter is 'a', which does not
 * match 'b'.
 *
 *
 * Example 4:
 *
 *
 * Input:
 * s = "adceb"
 * p = "*a*b"
 * Output: true
 * Explanation: The first '*' matches the empty sequence, while the second '*'
 * matches the substring "dce".
 *
 *
 * Example 5:
 *
 *
 * Input:
 * s = "acdcb"
 * p = "a*c?b"
 * Output: false
 *
 *
 */

// @lc code=start

/* 超时，需要转换成动归
func isMatch(s string, p string) bool {
	return matching(s, p, 0, 0)
}

func matching(s, p string, i, j int) bool {
	if i >= len(s) && j >= len(p) {
		return true
	}
	if i >= len(s) {
		if p[j] == '*' {
			return matching(s, p, i, j+1)
		} else {
			return false
		}
	}
	if j >= len(p) {
		return false
	}

	if s[i] == p[j] {
		return matching(s, p, i+1, j+1)
	} else if p[j] == '?' {
		return matching(s, p, i+1, j+1)
	} else if p[j] == '*' { // 匹配0次或多次
		// 尝试匹配0次
		zero := matching(s, p, i-1, j+1)
		if zero == true {
			return true
		} else { // 尝试匹配多次
			return matching(s, p, i+1, j)
		}
	}

	return false

}
*/

func isMatch(s string, p string) bool {
	slen, plen := len(s), len(p)
	var dp [][]bool
	for i := 0; i <= slen; i++ {
		t := make([]bool, plen+1)
		dp = append(dp, t)
	}
	dp[0][0] = true
	for i := 1; i <= plen; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i <= slen; i++ {
		for j := 1; j <= plen; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j] // *匹配空或者匹配一个
			}
		}
	}

	return dp[slen][plen]
}

// @lc code=end

