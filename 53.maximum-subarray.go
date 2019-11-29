import "math"

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (44.82%)
 * Likes:    5411
 * Dislikes: 224
 * Total Accepted:    675.5K
 * Total Submissions: 1.5M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 *
 * Example:
 *
 *
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 *
 * Follow up:
 *
 * If you have figured out the O(n) solution, try coding another solution using
 * the divide and conquer approach, which is more subtle.
 *
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last := nums[0]
	sum := []int{last}
	for i := 1; i < len(nums); i++ {
		if nums[i] < last+nums[i] {
			last += nums[i]
		} else {
			last = nums[i]
		}
		sum = append(sum, last)
	}

	res := math.MinInt64
	for i := 0; i < len(sum); i++ {
		if res < sum[i] {
			res = sum[i]
		}
	}
	return res
}

// @lc code=end

