/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 *
 * https://leetcode.com/problems/edit-distance/description/
 *
 * algorithms
 * Hard (39.79%)
 * Likes:    2677
 * Dislikes: 41
 * Total Accepted:    207.3K
 * Total Submissions: 516.1K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * Given two words word1 and word2, find the minimum number of operations
 * required to convert word1 to word2.
 *
 * You have the following 3 operations permitted on a word:
 *
 *
 * Insert a character
 * Delete a character
 * Replace a character
 *
 *
 * Example 1:
 *
 *
 * Input: word1 = "horse", word2 = "ros"
 * Output: 3
 * Explanation:
 * horse -> rorse (replace 'h' with 'r')
 * rorse -> rose (remove 'r')
 * rose -> ros (remove 'e')
 *
 *
 * Example 2:
 *
 *
 * Input: word1 = "intention", word2 = "execution"
 * Output: 5
 * Explanation:
 * intention -> inention (remove 't')
 * inention -> enention (replace 'i' with 'e')
 * enention -> exention (replace 'n' with 'x')
 * exention -> exection (replace 'n' with 'c')
 * exection -> execution (insert 'u')
 *
 *
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	if l1 == 0 {
		return l2
	}
	if l2 == 0 {
		return l1
	}

	dp := [][]int{}
	for i := 0; i <= l1; i++ {
		tmp := make([]int, l2+1)
		dp = append(dp, tmp)
	}
	dp[0][0] = 0
	for i := 1; i <= l1; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= l2; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}

	return dp[l1][l2]
}

func min(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}

// @lc code=end

