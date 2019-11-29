/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 *
 * https://leetcode.com/problems/subsets/description/
 *
 * algorithms
 * Medium (55.68%)
 * Likes:    2550
 * Dislikes: 62
 * Total Accepted:    443.4K
 * Total Submissions: 788.1K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a set of distinct integers, nums, return all possible subsets (the
 * power set).
 *
 * Note: The solution set must not contain duplicate subsets.
 *
 * Example:
 *
 *
 * Input: nums = [1,2,3]
 * Output:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 *
 */

// @lc code=start
func subsets(nums []int) [][]int {
	l := len(nums)
	result := [][]int{}
	for i := 0; i <= l; i++ {
		tmp := [][]int{}
		helper(&tmp, 0, i, 0, nums, []int{})
		result = append(result, tmp...)
	}
	return result
}

func helper(res *[][]int, level, k, i int, nums, r []int) {
	if level == k {
		tmp := make([]int, len(r))
		copy(tmp, r)
		*res = append(*res, tmp)
		return
	}

	for ; i < len(nums); i++ {
		r = append(r, nums[i])
		helper(res, level+1, k, i+1, nums, r)
		r = r[:len(r)-1]
	}
}

// @lc code=end

