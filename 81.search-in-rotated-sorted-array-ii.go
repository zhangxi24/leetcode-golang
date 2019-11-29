/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Medium (32.79%)
 * Likes:    867
 * Dislikes: 370
 * Total Accepted:    198.1K
 * Total Submissions: 603.3K
 * Testcase Example:  '[2,5,6,0,0,1,2]\n0'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 *
 * (i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).
 *
 * You are given a target value to search. If found in the array return true,
 * otherwise return false.
 *
 * Example 1:
 *
 *
 * Input: nums = [2,5,6,0,0,1,2], target = 0
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2,5,6,0,0,1,2], target = 3
 * Output: false
 *
 * Follow up:
 *
 *
 * This is a follow up problem to Search in Rotated Sorted Array, where nums
 * may contain duplicates.
 * Would this affect the run-time complexity? How and why?
 *
 *
 */

// @lc code=start
func search(nums []int, target int) bool {
	min, max := 0, len(nums)-1
	for min <= max {
		mid := (min + max) / 2
		if nums[mid] == target {
			return true
		} else if nums[min] == nums[mid] && nums[mid] == nums[max] {
			min++
			max--
		} else if nums[min] <= nums[mid] { // 左边升序
			if nums[min] <= target && nums[mid] > target {
				max = mid - 1
			} else {
				min = mid + 1
			}
		} else { // 右边升序
			if nums[mid] < target && nums[max] >= target {
				min = mid + 1
			} else {
				max = mid - 1
			}
		}
	}
	return false
}

// @lc code=end

