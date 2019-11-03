/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 *
 * https://leetcode.com/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (41.24%)
 * Likes:    1596
 * Dislikes: 207
 * Total Accepted:    471.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * Given a sorted array and a target value, return the index if the target is
 * found. If not, return the index where it would be if it were inserted in
 * order.
 *
 * You may assume no duplicates in the array.
 *
 * Example 1:
 *
 *
 * Input: [1,3,5,6], 5
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: [1,3,5,6], 2
 * Output: 1
 *
 *
 * Example 3:
 *
 *
 * Input: [1,3,5,6], 7
 * Output: 4
 *
 *
 * Example 4:
 *
 *
 * Input: [1,3,5,6], 0
 * Output: 0
 *
 *
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	min, max := 0, len(nums)-1
	pos := -1
	for min <= max {
		mid := (min + max) / 2
		if nums[mid] == target {
			pos = mid
			break
		} else if nums[mid] < target {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	if pos == -1 {
		if max == -1 {
			pos = 0
		} else if min > len(nums)-1 {
			pos = len(nums)
		} else {
			pos = min
		}
	}
	return pos
}

// @lc code=end

