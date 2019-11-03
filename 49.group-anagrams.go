/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 *
 * https://leetcode.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (50.00%)
 * Likes:    2172
 * Dislikes: 134
 * Total Accepted:    425.6K
 * Total Submissions: 844.6K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * Given an array of strings, group anagrams together.
 *
 * Example:
 *
 *
 * Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * Output:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 *
 * Note:
 *
 *
 * All inputs will be in lowercase.
 * The order of your output does not matter.
 *
 *
 */

// @lc code=start

// 字典排序
/*
func strSort(w string) string {
	ws := []string{}
	for i := 0; i < len(w); i++ {
		ws = append(ws, w[i:i+1])
	}
	sort.Strings(ws)
	res := ""
	for _, v := range ws {
		res += v
	}
	return res
}

func groupAnagrams(strs []string) [][]string {
	dict := map[string][]string{}
	for _, w := range strs {
		key := strSort(w)
		dict[key] = append(dict[key], w)
	}

	res := [][]string{}
	for _, v := range dict {
		res = append(res, v)
	}
	return res
}
*/

// 质数
func groupAnagrams(strs []string) [][]string {
	m := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	dict := map[int][]string{}
	for _, word := range strs {
		mul := 1
		for _, w := range word {
			mul *= m[int(w-'a')]
		}
		dict[mul] = append(dict[mul], word)
	}

	res := [][]string{}
	for _, v := range dict {
		res = append(res, v)
	}
	return res
}

// @lc code=end

