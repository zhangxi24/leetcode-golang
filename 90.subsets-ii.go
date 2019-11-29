import "sort"

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 *
 * https://leetcode.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (44.01%)
 * Likes:    1179
 * Dislikes: 55
 * Total Accepted:    235K
 * Total Submissions: 529.4K
 * Testcase Example:  '[1,2,2]'
 *
 * Given a collection of integers that might contain duplicates, nums, return
 * all possible subsets (the power set).
 *
 * Note: The solution set must not contain duplicate subsets.
 *
 * Example:
 *
 *
 * Input: [1,2,2]
 * Output:
 * [
 * ⁠ [2],
 * ⁠ [1],
 * ⁠ [1,2,2],
 * ⁠ [2,2],
 * ⁠ [1,2],
 * ⁠ []
 * ]
 *
 *
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	l := len(nums)
	result := [][]int{[]int{}}
	rlen := 0 // 上一次插入开始位置
	sort.Ints(nums)
	for i := 0; i < l; i++ {
		j := 0
		if i > 0 && nums[i] == nums[i-1] {
			j = rlen
		}
		rlen = len(result)
		for ; j < rlen; j++ {
			r := make([]int, len(result[j]))
			copy(r, result[j])
			r = append(r, nums[i])
			result = append(result, r)
		}
	}
	return result
}

// @lc code=end

