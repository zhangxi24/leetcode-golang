/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (28.99%)
 * Likes:    6735
 * Dislikes: 396
 * Total Accepted:    1.2M
 * Total Submissions: 4M
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string, find the length of the longest substring without repeating
 * characters.
 *
 *
 * Example 1:
 *
 *
 * Input: "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "bbbbb"
 * Output: 1
 * Explanation: The answer is "b", with the length of 1.
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "pwwkew"
 * Output: 3
 * Explanation: The answer is "wke", with the length of 3.
 * ⁠            Note that the answer must be a substring, "pwke" is a
 * subsequence and not a substring.
 *
 *
 *
 *
 *
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m := map[rune]int{}
	maxlen, l, begin := 0, 0, 0
	for i, ch := range s {
		if _, ok := m[ch]; ok {
			if begin < m[ch] {
				begin = m[ch]
			}
			l = i - begin
		} else {
			l += 1
		}
		if l > maxlen {
			maxlen = l
		}
		m[ch] = i
	}
	return maxlen
}

// @lc code=end

