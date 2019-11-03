/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 *
 * https://leetcode.com/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (33.56%)
 * Likes:    1236
 * Dislikes: 3735
 * Total Accepted:    374.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
 * of rows like this: (you may want to display this pattern in a fixed font for
 * better legibility)
 *
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 *
 * And then read line by line: "PAHNAPLSIIGYIR"
 *
 * Write the code that will take a string and make this conversion given a
 * number of rows:
 *
 *
 * string convert(string s, int numRows);
 *
 * Example 1:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 3
 * Output: "PAHNAPLSIIGYIR"
 *
 *
 * Example 2:
 *
 *
 * Input: s = "PAYPALISHIRING", numRows = 4
 * Output: "PINALSIGYAHRPI"
 * Explanation:
 *
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 *
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := (numRows - 1) * 2
	var sub []byte
	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			for a := 0; n*a+i < len(s); a++ {
				sub = append(sub, s[n*a+i])
			}
		} else {
			for a := 0; n*a+i < len(s) || n*a+n-i < len(s); a++ {
				if n*a+i < len(s) {
					sub = append(sub, s[n*a+i])
				}
				if n*a+n-i < len(s) {
					sub = append(sub, s[n*a+n-i])
				}
			}
		}
	}
	return string(sub[:])
}

// @lc code=end

