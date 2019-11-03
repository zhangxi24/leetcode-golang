import "sort"

/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 *
 * https://leetcode.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (43.92%)
 * Likes:    1149
 * Dislikes: 51
 * Total Accepted:    260.2K
 * Total Submissions: 590.2K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * Given a collection of candidate numbers (candidates) and a target number
 * (target), find all unique combinations in candidates where the candidate
 * numbers sums to target.
 *
 * Each number in candidates may only be used once in the combination.
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
 * Input: candidates = [10,1,2,7,6,1,5], target = 8,
 * A solution set is:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,5,2,1,2], target = 5,
 * A solution set is:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return helper(candidates, 0, target)
}

func helper(candidates []int, index, target int) [][]int {
	res := [][]int{}
	last := -1
	for i := index; i < len(candidates); i++ {
		if last == candidates[i] {
			continue
		}
		last = candidates[i]
		num := candidates[i]
		if num > target {
			break
		} else if num == target {
			res = append(res, []int{num})
		} else {
			r := helper(candidates, i+1, target-num)
			for _, subset := range r {
				s := []int{num}
				res = append(res, append(s, subset...))
			}
		}
	}

	return res
}

// @lc code=end

