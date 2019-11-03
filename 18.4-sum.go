/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 *
 * https://leetcode.com/problems/4sum/description/
 *
 * algorithms
 * Medium (31.68%)
 * Likes:    1316
 * Dislikes: 265
 * Total Accepted:    271.4K
 * Total Submissions: 854.7K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * Given an array nums of n integers and an integer target, are there elements
 * a, b, c, and d in nums such that a + b + c + d = target? Find all unique
 * quadruplets in the array which gives the sum of target.
 * 
 * Note:
 * 
 * The solution set must not contain duplicate quadruplets.
 * 
 * Example:
 * 
 * 
 * Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
 * 
 * A solution set is:
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 * 
 * 
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			t := target - nums[i] - nums[j]
			k, l := j+1, len(nums)-1
			for k < l {
				if nums[k]+nums[l] == t {
					a, b := nums[k], nums[l]
					ans := []int{nums[i], nums[j], nums[k], nums[l]}
					res = append(res, ans)
					k++
					l--
					for k < l && a == nums[k] {
						k++
					}
					for k < l && b == nums[l] {
						l--
					}
				} else if nums[k]+nums[l] > t && k < l {
					l--
				} else if nums[k]+nums[l] < t && k < l {
					k++
				}
			}
			for j+1 < len(nums)-2 && nums[j+1] == nums[j] {
				j++
			}
		}
		for i+1 < len(nums)-3 && nums[i+1] == nums[i] {
			i++
		}
	}

	return res
}
// @lc code=end

