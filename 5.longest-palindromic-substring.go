/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (28.08%)
 * Likes:    4610
 * Dislikes: 417
 * Total Accepted:    694K
 * Total Submissions: 2.5M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, find the longest palindromic substring in s. You may
 * assume that the maximum length of s is 1000.
 *
 * Example 1:
 *
 *
 * Input: "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 *
 *
 * Example 2:
 *
 *
 * Input: "cbbd"
 * Output: "bb"
 *
 *
 */

// @lc code=start
func longestPalindrome(s string) string {
	maxlen := 0
	maxstr := ""
	for i := 0; i < len(s); i++ {
		len, str := iCenterPalindrome(s, i)
		if len > maxlen {
			maxlen = len
			maxstr = str
		}
	}
	return maxstr
}

func iCenterPalindrome(s string, i int) (maxlen int, maxstr string) {
	maxstr = s[i : i+1]
	maxlen = 1
	// i为中心
	for j := 1; i-j >= 0 && i+j < len(s); j++ {
		if s[i-j] != s[i+j] {
			break
		}
		maxstr = s[i-j : i+j+1]
		maxlen += 2
	}
	// (i+1)/2为中心
	tmplen := 0
	for j := 0; i-j >= 0 && i+1+j < len(s); j++ {
		if s[i-j] != s[i+j+1] {
			break
		}
		tmplen += 2
		if tmplen > maxlen {
			maxlen = tmplen
			maxstr = s[i-j : i+j+2]
		}
	}

	return
}

// @lc code=end

