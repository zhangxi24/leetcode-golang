/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 *
 * https://leetcode.com/problems/substring-with-concatenation-of-all-words/description/
 *
 * algorithms
 * Hard (24.25%)
 * Likes:    646
 * Dislikes: 1027
 * Total Accepted:    149.3K
 * Total Submissions: 615.2K
 * Testcase Example:  '"barfoothefoobarman"\n["foo","bar"]'
 *
 * You are given a string, s, and a list of words, words, that are all of the
 * same length. Find all starting indices of substring(s) in s that is a
 * concatenation of each word in words exactly once and without any intervening
 * characters.
 *
 *
 *
 * Example 1:
 *
 *
 * Input:
 * ⁠ s = "barfoothefoobarman",
 * ⁠ words = ["foo","bar"]
 * Output: [0,9]
 * Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar"
 * respectively.
 * The output order does not matter, returning [9,0] is fine too.
 *
 *
 * Example 2:
 *
 *
 * Input:
 * ⁠ s = "wordgoodgoodgoodbestword",
 * ⁠ words = ["word","good","best","word"]
 * Output: []
 *
 *
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	pos := []int{}
	wordnum := len(words)
	if wordnum == 0 {
		return pos
	}
	wordlen := len(words[0])

	wordmap := map[string]int{}
	for _, w := range words {
		if _, ok := wordmap[w]; ok {
			wordmap[w]++
		} else {
			wordmap[w] = 1
		}
	}

	for i := 0; i+wordlen <= len(s); i++ {
		exist := map[string]int{}
		p := i
		for n := 0; n < wordnum && p+wordlen <= len(s); n++ {
			e := 0
			for k, _ := range wordmap {
				if s[p:p+wordlen] == k {
					e = 1
					if _, ok := exist[k]; ok {
						exist[k]++
					} else {
						exist[k] = 1
					}
				}
			}
			if e == 0 {
				break
			}
			p += wordlen
		}
		if mapCompare(exist, wordmap) == true {
			pos = append(pos, i)
		}
	}

	return pos
}

func mapCompare(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range b {
		if num, ok := a[k]; !ok || num != v {
			return false
		}
	}
	return true
}

// @lc code=end

