/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 *
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (32.32%)
 * Likes:    466
 * Dislikes: 1904
 * Total Accepted:    309.1K
 * Total Submissions: 956.1K
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consists of upper/lower-case alphabets and empty space
 * characters ' ', return the length of last word in the string.
 *
 * If the last word does not exist, return 0.
 *
 * Note: A word is defined as a character sequence consists of non-space
 * characters only.
 *
 * Example:
 *
 *
 * Input: "Hello World"
 * Output: 5
 *
 *
 *
 *
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	l, last := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if l != 0 {
				last = l
			}
			l = 0
		} else {
			l++
			if i == len(s)-1 {
				last = l
			}
		}
	}
	return last
}

/*
// strings.Trim:去掉字符串s中首部以及尾部与字符串cutset中每个相匹配的字符
func lengthOfLastWord(s string) int {
    s = strings.Trim(s, " ")
    if len(s)==0 {
        return 0
	}
    sArray := strings.Split(s, " ")
    return len(sArray[len(sArray)-1])
}
*/

// @lc code=end

