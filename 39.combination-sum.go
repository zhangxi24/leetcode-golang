/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 *
 * https://leetcode.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (51.27%)
 * Likes:    2598
 * Dislikes: 79
 * Total Accepted:    414.2K
 * Total Submissions: 804.7K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * Given a set of candidate numbers (candidates) (without duplicates) and a
 * target number (target), find all unique combinations in candidates where the
 * candidate numbers sums to target.
 *
 * The same repeated number may be chosen from candidates unlimited number of
 * times.
 *
 * Note:
 *
 *
 * All numbers (including target) will be positive integers.
 * The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 *
 * Input: candidates = [2,3,6,7], target = 7,
 * A solution set is:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,3,5], target = 8,
 * A solution set is:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 *
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	return helper(candidates, target)
}

func helper(candidates []int, target int) [][]int {
	res := [][]int{}
	for n, num := range candidates {
		adder := []int{}
		for dnum := num; dnum <= target; dnum += num {
			adder = append(adder, num)
			if dnum == target {
				res = append(res, adder)
			} else if n != len(candidates)-1 { // 其他加数构成target-dnum
				can := (candidates)[n+1:]
				r := helper(can, target-dnum)
				for _, subset := range r {
					// 不能用append(adder, subset...)
					d := make([]int, len(adder))
					copy(d, adder)
					res = append(res, append(d, subset...))
				}
			}
		}
	}
	return res
}

// @lc code=end

