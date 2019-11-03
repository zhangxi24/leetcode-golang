import "sort"

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (24.97%)
 * Likes:    4753
 * Dislikes: 567
 * Total Accepted:    680.7K
 * Total Submissions: 2.7M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an array nums of n integers, are there elements a, b, c in nums such
 * that a + b + c = 0? Find all unique triplets in the array which gives the
 * sum of zero.
 *
 * Note:
 *
 * The solution set must not contain duplicate triplets.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 0, 1, 2, -1, -4],
 *
 * A solution set is:
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 */

// @lc code=start

func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for min := 0; nums[min] <= 0 && min < len(nums)-2; min++ {
		if min > 0 && nums[min] == nums[min-1] {
			continue
		}
		for max := len(nums) - 1; nums[max] >= 0 && max > min+1; max-- {
			if max != len(nums)-1 && nums[max] == nums[max+1] {
				continue
			}
			target := 0 - nums[min] - nums[max]
			if target > nums[max] || target < nums[min] {
				continue
			}
			if target > 0 {
				mid := max - 1
				for nums[mid] > target {
					mid--
				}
				if nums[mid] == target {
					res = append(res, []int{nums[min], nums[mid], nums[max]})
				}
			} else {
				mid := min + 1
				for nums[mid] < target {
					mid++
				}
				if nums[mid] == target {
					res = append(res, []int{nums[min], nums[mid], nums[max]})
				}
			}
		}
	}

	return res
}

// @lc code=end

