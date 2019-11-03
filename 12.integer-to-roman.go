/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 *
 * https://leetcode.com/problems/integer-to-roman/description/
 *
 * algorithms
 * Medium (52.37%)
 * Likes:    728
 * Dislikes: 2199
 * Total Accepted:    275.6K
 * Total Submissions: 525.3K
 * Testcase Example:  '3'
 *
 * Roman numerals are represented by seven different symbols: I, V, X, L, C, D
 * and M.
 * 
 * 
 * Symbol       Value
 * I             1
 * V             5
 * X             10
 * L             50
 * C             100
 * D             500
 * M             1000
 * 
 * For example, two is written as II in Roman numeral, just two one's added
 * together. Twelve is written as, XII, which is simply X + II. The number
 * twenty seven is written as XXVII, which is XX + V + II.
 * 
 * Roman numerals are usually written largest to smallest from left to right.
 * However, the numeral for four is not IIII. Instead, the number four is
 * written as IV. Because the one is before the five we subtract it making
 * four. The same principle applies to the number nine, which is written as IX.
 * There are six instances where subtraction is used:
 * 
 * 
 * I can be placed before V (5) and X (10) to make 4 and 9. 
 * X can be placed before L (50) and C (100) to make 40 and 90. 
 * C can be placed before D (500) and M (1000) to make 400 and 900.
 * 
 * 
 * Given an integer, convert it to a roman numeral. Input is guaranteed to be
 * within the range from 1 to 3999.
 * 
 * Example 1:
 * 
 * 
 * Input: 3
 * Output: "III"
 * 
 * Example 2:
 * 
 * 
 * Input: 4
 * Output: "IV"
 * 
 * Example 3:
 * 
 * 
 * Input: 9
 * Output: "IX"
 * 
 * Example 4:
 * 
 * 
 * Input: 58
 * Output: "LVIII"
 * Explanation: L = 50, V = 5, III = 3.
 * 
 * 
 * Example 5:
 * 
 * 
 * Input: 1994
 * Output: "MCMXCIV"
 * Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
 * 
 */

// @lc code=start
func intToRoman(num int) string {
	res := ""

	thousand := num / 1000
	for i := 0; i < thousand; i++ {
		res += "M"
	}

	num = num % 1000
	hundred := num / 100
	if hundred == 9 {
		res += "CM"
	} else if hundred >= 5 && hundred < 9 {
		res += "D"
		for i := 0; i < hundred-5; i++ {
			res += "C"
		}
	} else if hundred == 4 {
		res += "CD"
	} else {
		for i := 0; i < hundred; i++ {
			res += "C"
		}
	}

	num = num % 100
	ten := num / 10
	if ten == 9 {
		res += "XC"
	} else if ten >= 5 && ten < 9 {
		res += "L"
		for i := 0; i < ten-5; i++ {
			res += "X"
		}
	} else if ten == 4 {
		res += "XL"
	} else {
		for i := 0; i < ten; i++ {
			res += "X"
		}
	}

	num = num % 10
	if num == 9 {
		res += "IX"
	} else if num >= 5 && num < 9 {
		res += "V"
		for i := 0; i < num-5; i++ {
			res += "I"
		}
	} else if num == 4 {
		res += "IV"
	} else {
		for i := 0; i < num; i++ {
			res += "I"
		}
	}

	return res
}
// @lc code=end

