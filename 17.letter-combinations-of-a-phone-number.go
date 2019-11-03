/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (43.35%)
 * Likes:    2720
 * Dislikes: 343
 * Total Accepted:    466.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '"23"'
 *
 * Given a string containing digits from 2-9 inclusive, return all possible
 * letter combinations that the number could represent.
 *
 * A mapping of digit to letters (just like on the telephone buttons) is given
 * below. Note that 1 does not map to any letters.
 *
 *
 *
 * Example:
 *
 *
 * Input: "23"
 * Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * Note:
 *
 * Although the above answer is in lexicographical order, your answer could be
 * in any order you want.
 *
 */

// @lc code=start
func letterCombinations(digits string) []string {
	LetterMap := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	res := []string{}
	for i := 0; i < len(digits); i++ {
		if _, ok := LetterMap[digits[i:i+1]]; ok {
			add := LetterMap[digits[i:i+1]]
			if i == 0 {
				res = add
			} else {
				newres := []string{}
				for _, r := range res {
					for _, l := range add {
						newres = append(newres, r+l)
					}
				}
				res = newres
			}
		}
	}
	return res
}

// @lc code=end

