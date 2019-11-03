/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (34.24%)
 * Likes:    1684
 * Dislikes: 1520
 * Total Accepted:    569K
 * Total Submissions: 1.7M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 * Example 1:
 *
 *
 * Input: ["flower","flow","flight"]
 * Output: "fl"
 *
 *
 * Example 2:
 *
 *
 * Input: ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 * Note:
 *
 * All given inputs are in lowercase letters a-z.
 *
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs) == 0 {
		return prefix
	} else {
		prefix = strs[0]
	}
	for i, str := range strs {
		if str == "" {
			return ""
		}
		if i == 0 {
			continue
		}
		j, cut := 0, false
		for ; j < len(str) && j < len(prefix); j++ {
			if str[j] != prefix[j] {
				prefix = prefix[:j]
				cut = true
				break
			}
		}
		if cut == false && j == len(str) {
			prefix = str
		}
	}
	return prefix
}

// @lc code=end

