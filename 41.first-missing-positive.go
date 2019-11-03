/*
 * @lc app=leetcode id=41 lang=golang
 *
 * [41] First Missing Positive
 *
 * https://leetcode.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (29.97%)
 * Likes:    2238
 * Dislikes: 661
 * Total Accepted:    254.8K
 * Total Submissions: 847.5K
 * Testcase Example:  '[1,2,0]'
 *
 * Given an unsorted integer array, find the smallest missing positive
 * integer.
 *
 * Example 1:
 *
 *
 * Input: [1,2,0]
 * Output: 3
 *
 *
 * Example 2:
 *
 *
 * Input: [3,4,-1,1]
 * Output: 2
 *
 *
 * Example 3:
 *
 *
 * Input: [7,8,9,11,12]
 * Output: 1
 *
 *
 * Note:
 *
 * Your algorithm should run in O(n) time and uses constant extra space.
 *
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	for i := range nums {
		j := i
		for {
			if j < 0 || j >= len(nums) {
				break
			}
			if nums[j] <= 0 || nums[j] > len(nums) || nums[j] == j+1 {
				break
			}
			if nums[j] == nums[nums[j]-1] {
				nums[j] = 0
				break
			}
			nums[j], nums[nums[j]-1] = nums[nums[j]-1], nums[j]
		}
	}

	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

// @lc code=end

