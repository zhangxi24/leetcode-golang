import "strconv"

/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 *
 * https://leetcode.com/problems/count-and-say/description/
 *
 * algorithms
 * Easy (42.00%)
 * Likes:    928
 * Dislikes: 7391
 * Total Accepted:    325.3K
 * Total Submissions: 773.2K
 * Testcase Example:  '1'
 *
 * The count-and-say sequence is the sequence of integers with the first five
 * terms as following:
 *
 *
 * 1.     1
 * 2.     11
 * 3.     21
 * 4.     1211
 * 5.     111221
 *
 *
 * 1 is read off as "one 1" or 11.
 * 11 is read off as "two 1s" or 21.
 * 21 is read off as "one 2, then one 1" or 1211.
 *
 * Given an integer n where 1 ≤ n ≤ 30, generate the n^th term of the
 * count-and-say sequence.
 *
 * Note: Each term of the sequence of integers will be represented as a
 * string.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: 1
 * Output: "1"
 *
 *
 * Example 2:
 *
 *
 * Input: 4
 * Output: "1211"
 *
 */

// @lc code=start
func countAndSay(n int) string {
	res := ""
	if n == 1 {
		res = "1"
	} else {
		pre := countAndSay(n - 1)
		for i := 0; i < len(pre); {
			n := pre[i : i+1]
			num := 1
			i++
			for ; i < len(pre) && pre[i:i+1] == n; i++ {
				num++
			}
			add := strconv.Itoa(num) + n
			res += add
		}
	}
	return res
}

// @lc code=end

