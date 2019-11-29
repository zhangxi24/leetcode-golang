/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 *
 * https://leetcode.com/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (32.19%)
 * Likes:    3053
 * Dislikes: 218
 * Total Accepted:    298.2K
 * Total Submissions: 917.5K
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * Given a string S and a string T, find the minimum window in S which will
 * contain all the characters in T in complexity O(n).
 *
 * Example:
 *
 *
 * Input: S = "ADOBECODEBANC", T = "ABC"
 * Output: "BANC"
 *
 *
 * Note:
 *
 *
 * If there is no such window in S that covers all characters in T, return the
 * empty string "".
 * If there is such window, you are guaranteed that there will always be only
 * one unique minimum window in S.
 *
 *
 */

// @lc code=start
func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	need := map[string]int{}
	for i := 0; i < len(t); i++ {
		if _, ok := need[t[i:i+1]]; ok {
			need[t[i:i+1]]++
		} else {
			need[t[i:i+1]] = 1
		}
	}

	window := map[string]int{}
	l := len(s) + 1
	left := 0
	res := ""
	for right := 0; right < len(s); right++ {
		if _, ok := window[s[right:right+1]]; ok {
			window[s[right:right+1]]++
		} else {
			window[s[right:right+1]] = 1
		}
		for checkWindow(window, need) && left <= right {
			if right+1-left < l {
				l = right + 1 - left
				res = s[left : right+1]
			}
			window[s[left:left+1]]--
			left++
		}
	}
	return res
}

func checkWindow(window, need map[string]int) bool {
	for k, v := range need {
		if d, ok := window[k]; !ok || v > d {
			return false
		}
	}
	return true
}

// @lc code=end

