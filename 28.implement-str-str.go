/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 *
 * https://leetcode.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (33.08%)
 * Likes:    1080
 * Dislikes: 1519
 * Total Accepted:    509.3K
 * Total Submissions: 1.5M
 * Testcase Example:  '"hello"\n"ll"'
 *
 * Implement strStr().
 *
 * Return the index of the first occurrence of needle in haystack, or -1 if
 * needle is not part of haystack.
 *
 * Example 1:
 *
 *
 * Input: haystack = "hello", needle = "ll"
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: haystack = "aaaaa", needle = "bba"
 * Output: -1
 *
 *
 * Clarification:
 *
 * What should we return when needle is an empty string? This is a great
 * question to ask during an interview.
 *
 * For the purpose of this problem, we will return 0 when needle is an empty
 * string. This is consistent to C's strstr() and Java's indexOf().
 *
 */

// @lc code=start
func getNextVal(needle string) []int {
	nextval := []int{}
	k := -1
	for j := -1; j < len(needle)-1; {
		if j == -1 {
			nextval = append(nextval, -1)
			j++
		} else if k == -1 || needle[j] == needle[k] {
			j++
			k++
			if needle[k] != needle[j] {
				nextval = append(nextval, k)
			} else {
				nextval = append(nextval, nextval[k])
			}
		} else {
			k = nextval[k]
		}
	}
	return nextval
}

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 && len(needle) == 0 {
		return 0
	}
	nextval := getNextVal(needle)
	i := 0
	j := 0
	for i < len(haystack) && j < len(needle) {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = nextval[j]
		}
	}
	if j >= len(needle) {
		return i - len(needle)
	}
	return -1
}

// @lc code=end

