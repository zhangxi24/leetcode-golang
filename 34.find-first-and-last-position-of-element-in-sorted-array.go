/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 *
 * https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (34.39%)
 * Likes:    2139
 * Dislikes: 100
 * Total Accepted:    367.9K
 * Total Submissions: 1.1M
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * Given an array of integers nums sorted in ascending order, find the starting
 * and ending position of a given target value.
 *
 * Your algorithm's runtime complexity must be in the order of O(log n).
 *
 * If the target is not found in the array, return [-1, -1].
 *
 * Example 1:
 *
 *
 * Input: nums = [5,7,7,8,8,10], target = 8
 * Output: [3,4]
 *
 * Example 2:
 *
 *
 * Input: nums = [5,7,7,8,8,10], target = 6
 * Output: [-1,-1]
 *
 */

// @lc code=start

func searchRange(nums []int, target int) []int {
	pos := -1
	min, max := 0, len(nums)-1
	res := []int{}
	for min <= max {
		mid := (min + max) / 2
		if nums[mid] == target {
			pos = mid
			break
		} else if target > nums[mid] {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	if pos == -1 {
		res = append(res, pos)
		res = append(res, pos)
	} else {
		low, high := pos, pos
		for ; low >= 0 && nums[low] == target; low-- {
		}
		res = append(res, low+1)
		for ; high < len(nums) && nums[high] == target; high++ {
		}
		res = append(res, high-1)
	}
	return res
}

// @lc code=end

