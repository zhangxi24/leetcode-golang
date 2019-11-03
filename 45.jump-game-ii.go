/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 *
 * https://leetcode.com/problems/jump-game-ii/description/
 *
 * algorithms
 * Hard (28.98%)
 * Likes:    1573
 * Dislikes: 87
 * Total Accepted:    202.6K
 * Total Submissions: 695.6K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * Given an array of non-negative integers, you are initially positioned at the
 * first index of the array.
 *
 * Each element in the array represents your maximum jump length at that
 * position.
 *
 * Your goal is to reach the last index in the minimum number of jumps.
 *
 * Example:
 *
 *
 * Input: [2,3,1,1,4]
 * Output: 2
 * Explanation: The minimum number of jumps to reach the last index is 2.
 * ⁠   Jump 1 step from index 0 to 1, then 3 steps to the last index.
 *
 * Note:
 *
 * You can assume that you can always reach the last index.
 *
 */

// @lc code=start
func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	step := 0
	start, end := 0, 0
	for end < len(nums)-1 {
		max := end
		for i := start; i <= end; i++ {
			if nums[i]+i > max {
				max = nums[i] + i
			}
		}
		start = end + 1
		end = max
		step++
	}

	return step
}

// @lc code=end

