/*
 * @lc app=leetcode id=87 lang=golang
 *
 * [87] Scramble String
 *
 * https://leetcode.com/problems/scramble-string/description/
 *
 * algorithms
 * Hard (32.37%)
 * Likes:    363
 * Dislikes: 580
 * Total Accepted:    98.5K
 * Total Submissions: 303.1K
 * Testcase Example:  '"great"\n"rgeat"'
 *
 * Given a string s1, we may represent it as a binary tree by partitioning it
 * to two non-empty substrings recursively.
 *
 * Below is one possible representation of s1 = "great":
 *
 *
 * ⁠   great
 * ⁠  /    \
 * ⁠ gr    eat
 * ⁠/ \    /  \
 * g   r  e   at
 * ⁠          / \
 * ⁠         a   t
 *
 *
 * To scramble the string, we may choose any non-leaf node and swap its two
 * children.
 *
 * For example, if we choose the node "gr" and swap its two children, it
 * produces a scrambled string "rgeat".
 *
 *
 * ⁠   rgeat
 * ⁠  /    \
 * ⁠ rg    eat
 * ⁠/ \    /  \
 * r   g  e   at
 * ⁠          / \
 * ⁠         a   t
 *
 *
 * We say that "rgeat" is a scrambled string of "great".
 *
 * Similarly, if we continue to swap the children of nodes "eat" and "at", it
 * produces a scrambled string "rgtae".
 *
 *
 * ⁠   rgtae
 * ⁠  /    \
 * ⁠ rg    tae
 * ⁠/ \    /  \
 * r   g  ta  e
 * ⁠      / \
 * ⁠     t   a
 *
 *
 * We say that "rgtae" is a scrambled string of "great".
 *
 * Given two strings s1 and s2 of the same length, determine if s2 is a
 * scrambled string of s1.
 *
 * Example 1:
 *
 *
 * Input: s1 = "great", s2 = "rgeat"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s1 = "abcde", s2 = "caebd"
 * Output: false
 *
 */

// @lc code=start
func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}

	var nums [26]int
	for i := 0; i < len(s1); i++ {
		nums[s1[i]-'a']++
	}
	for i := 0; i < len(s2); i++ {
		c := s2[i] - 'a'
		nums[c]--
		if nums[c] < 0 {
			return false
		}
	}

	for i := 1; i < len(s1); i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if isScramble(s1[:i], s2[len(s2)-i:]) && isScramble(s1[i:], s2[:len(s2)-i]) {
			return true
		}
	}
	return false
}

// @lc code=end

