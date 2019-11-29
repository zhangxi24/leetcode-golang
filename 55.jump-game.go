/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 *
 * https://leetcode.com/problems/jump-game/description/
 *
 * algorithms
 * Medium (32.70%)
 * Likes:    2555
 * Dislikes: 243
 * Total Accepted:    322.5K
 * Total Submissions: 983.9K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * Given an array of non-negative integers, you are initially positioned at the
 * first index of the array.
 *
 * Each element in the array represents your maximum jump length at that
 * position.
 *
 * Determine if you are able to reach the last index.
 *
 * Example 1:
 *
 *
 * Input: [2,3,1,1,4]
 * Output: true
 * Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last
 * index.
 *
 *
 * Example 2:
 *
 *
 * Input: [3,2,1,0,4]
 * Output: false
 * Explanation: You will always arrive at index 3 no matter what. Its
 * maximum
 * jump length is 0, which makes it impossible to reach the last index.
 *
 *
 */

// @lc code=start
func canJump(nums []int) bool {
	start, end := 0, 0
	for end < len(nums)-1 {
		min, max := end, end
		for i := start; i <= end; i++ {
			if max < nums[i]+i {
				max = nums[i] + i
			}
			if min > nums[i]+i && nums[i]+i > end {
				min = nums[i] + i
			}
		}
		start = min
		end = max
		if start == end {
			return false
		}
	}
	return true
}

// @lc code=end

