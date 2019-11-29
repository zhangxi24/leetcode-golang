/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 *
 * https://leetcode.com/problems/combinations/description/
 *
 * algorithms
 * Medium (50.25%)
 * Likes:    1018
 * Dislikes: 58
 * Total Accepted:    238.6K
 * Total Submissions: 470.8K
 * Testcase Example:  '4\n2'
 *
 * Given two integers n and k, return all possible combinations of k numbers
 * out of 1 ... n.
 *
 * Example:
 *
 *
 * Input: n = 4, k = 2
 * Output:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 *
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := [][]int{}
	helper(&res, 0, k, 1, n, map[int]bool{}, []int{})
	return res
}

func helper(res *[][]int, level, k, i, n int, used map[int]bool, r []int) {
	if level == k {
		tmp := make([]int, len(r))
		copy(tmp, r)
		*res = append(*res, tmp)
		return
	}

	for ; i <= n; i++ {
		if used[i] == false {
			used[i] = true
			r = append(r, i)
			helper(res, level+1, k, i+1, n, used, r)
			r = r[:len(r)-1]
			used[i] = false
		}
	}
}

// @lc code=end

