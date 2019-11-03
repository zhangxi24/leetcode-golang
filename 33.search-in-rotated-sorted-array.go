/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (33.26%)
 * Likes:    3128
 * Dislikes: 371
 * Total Accepted:    503.5K
 * Total Submissions: 1.5M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 *
 * (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
 *
 * You are given a target value to search. If found in the array return its
 * index, otherwise return -1.
 *
 * You may assume no duplicate exists in the array.
 *
 * Your algorithm's runtime complexity must be in the order of O(log n).
 *
 * Example 1:
 *
 *
 * Input: nums = [4,5,6,7,0,1,2], target = 0
 * Output: 4
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [4,5,6,7,0,1,2], target = 3
 * Output: -1
 *
 */

// @lc code=start
func search(nums []int, target int) int {
	pos := -1
	min, max := 0, len(nums)-1
	for min <= max {
		mid := (min + max) / 2
		if nums[min] <= nums[max] { // 此段顺序
			if target < nums[mid] {
				max = mid - 1
			} else if target > nums[mid] {
				min = mid + 1
			} else {
				pos = mid
				break
			}
		} else { // 存在逆序段
			if target == nums[mid] {
				pos = mid
				break
			} else if target < nums[min] && target > nums[max] {
				break
			} else if nums[min] <= nums[mid] { // 前一段顺序，后一段存在逆序
				if nums[min] <= target && target < nums[mid] { // 在顺序段中
					max = mid
				} else {
					min = mid + 1
				}
			} else { // 前一段存在逆序，后一段顺序
				if nums[mid] < target && target <= nums[max] { // 在顺序段中
					min = mid + 1
				} else {
					max = mid
				}
			}
		}
	}

	return pos
}

// @lc code=end

