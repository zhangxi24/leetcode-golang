/*
 * @lc app=leetcode id=97 lang=golang
 *
 * [97] Interleaving String
 *
 * https://leetcode.com/problems/interleaving-string/description/
 *
 * algorithms
 * Hard (29.18%)
 * Likes:    1011
 * Dislikes: 58
 * Total Accepted:    128.7K
 * Total Submissions: 436.5K
 * Testcase Example:  '"aabcc"\n"dbbca"\n"aadbbcbcac"'
 *
 * Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and
 * s2.
 *
 * Example 1:
 *
 *
 * Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
 * Output: false
 *
 *
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s1) == 0 || len(s2) == 0 {
		if s1 == s3 || s2 == s3 {
			return true
		} else {
			return false
		}
	}

	l1, l2 := len(s1), len(s2)
	dp := [][]bool{}
	for i := 0; i <= l1; i++ {
		tmp := make([]bool, l2+1)
		dp = append(dp, tmp)
	}
	dp[0][0] = true
	for i := 1; i <= l1; i++ {
		if s1[i-1] != s3[i-1] {
			break
		} else {
			dp[i][0] = true
		}
	}
	for i := 1; i <= l2; i++ {
		if s2[i-1] != s3[i-1] {
			break
		} else {
			dp[0][i] = true
		}
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if dp[i-1][j] == true && s1[i-1] == s3[i+j-1] {
				dp[i][j] = true
			} else if dp[i][j-1] == true && s2[j-1] == s3[i+j-1] {
				dp[i][j] = true
			}
		}
	}

	return dp[l1][l2]
}

// @lc code=end

