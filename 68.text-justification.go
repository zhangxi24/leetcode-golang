/*
 * @lc app=leetcode id=68 lang=golang
 *
 * [68] Text Justification
 *
 * https://leetcode.com/problems/text-justification/description/
 *
 * algorithms
 * Hard (24.67%)
 * Likes:    451
 * Dislikes: 1212
 * Total Accepted:    110.8K
 * Total Submissions: 445.1K
 * Testcase Example:  '["This", "is", "an", "example", "of", "text", "justification."]\n16'
 *
 * Given an array of words and a width maxWidth, format the text such that each
 * line has exactly maxWidth characters and is fully (left and right)
 * justified.
 *
 * You should pack your words in a greedy approach; that is, pack as many words
 * as you can in each line. Pad extra spaces ' ' when necessary so that each
 * line has exactly maxWidth characters.
 *
 * Extra spaces between words should be distributed as evenly as possible. If
 * the number of spaces on a line do not divide evenly between words, the empty
 * slots on the left will be assigned more spaces than the slots on the right.
 *
 * For the last line of text, it should be left justified and no extra space is
 * inserted between words.
 *
 * Note:
 *
 *
 * A word is defined as a character sequence consisting of non-space characters
 * only.
 * Each word's length is guaranteed to be greater than 0 and not exceed
 * maxWidth.
 * The input array words contains at least one word.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * words = ["This", "is", "an", "example", "of", "text", "justification."]
 * maxWidth = 16
 * Output:
 * [
 * "This    is    an",
 * "example  of text",
 * "justification.  "
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input:
 * words = ["What","must","be","acknowledgment","shall","be"]
 * maxWidth = 16
 * Output:
 * [
 * "What   must   be",
 * "acknowledgment  ",
 * "shall be        "
 * ]
 * Explanation: Note that the last line is "shall be    " instead of "shall
 * be",
 * because the last line must be left-justified instead of fully-justified.
 * ⁠            Note that the second line is also left-justified becase it
 * contains only one word.
 *
 *
 * Example 3:
 *
 *
 * Input:
 * words =
 * ["Science","is","what","we","understand","well","enough","to","explain",
 * "to","a","computer.","Art","is","everything","else","we","do"]
 * maxWidth = 20
 * Output:
 * [
 * "Science  is  what we",
 * ⁠ "understand      well",
 * "enough to explain to",
 * "a  computer.  Art is",
 * "everything  else  we",
 * "do                  "
 * ]
 *
 *
 */

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	num := len(words)
	if num == 0 {
		return words
	}

	res := []string{}
	pos := [][]int{}
	l := len(words[0])
	line := []int{0}
	if num == 1 {
		pos = append(pos, line)
	}
	for i := 1; i < num; i++ {
		if len(words[i])+1+l > maxWidth {
			tmp := make([]int, len(line))
			copy(tmp, line)
			pos = append(pos, tmp)
			line = line[:0]

			line = append(line, i)
			l = len(words[i])
		} else {
			l = l + len(words[i]) + 1
			line = append(line, i)
		}
		if i == num-1 {
			tmp := make([]int, len(line))
			copy(tmp, line)
			pos = append(pos, tmp)
			line = line[:0]
		}

	}

	for i := 0; i < len(pos)-1; i++ {
		spacenum := maxWidth
		for j := 0; j < len(pos[i]); j++ {
			spacenum -= len(words[pos[i][j]])
		}
		if len(pos[i]) == 1 {
			r := words[pos[i][0]]
			for j := len(words[pos[i][0]]); j < maxWidth; j++ {
				r += " "
			}
			res = append(res, r)
		} else {
			s := spacenum / (len(pos[i]) - 1)
			n := spacenum % (len(pos[i]) - 1)
			r := words[pos[i][0]]
			space := ""
			for m := 0; m < s; m++ {
				space += " "
			}
			for j := 1; j < len(pos[i]); j++ {
				r += space
				if j <= n {
					r += " "
				}
				r += words[pos[i][j]]
			}
			res = append(res, r)
		}
	}
	// 最后一行
	s := words[pos[len(pos)-1][0]]
	l = len(words[pos[len(pos)-1][0]])
	for i := 1; i < len(pos[len(pos)-1]); i++ {
		s = s + " " + words[pos[len(pos)-1][i]]
		l = l + 1 + len(words[pos[len(pos)-1][i]])
	}
	for j := l; j < maxWidth; j++ {
		s += " "
	}
	res = append(res, s)

	return res
}

// @lc code=end

