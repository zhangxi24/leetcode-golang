import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 *
 * https://leetcode.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (45.73%)
 * Likes:    1412
 * Dislikes: 103
 * Total Accepted:    392.3K
 * Total Submissions: 858K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * Given an array nums of n integers and an integer target, find three integers
 * in nums such that the sum is closest to target. Return the sum of the three
 * integers. You may assume that each input would have exactly one solution.
 *
 * Example:
 *
 *
 * Given array nums = [-1, 2, 1, -4], and target = 1.
 *
 * The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 *
 *
 */

// @lc code=start
func dis(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}

func threeSumClosest(nums []int, target int) int {
	mindis := math.MaxInt32
	sum := 0
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	for min := 0; min < len(nums)-2; min++ {
		t := target - nums[min]
		mid := min + 1
		max := len(nums) - 1
		for mid < max {
			d := dis(t, nums[mid]+nums[max])
			if d < mindis {
				mindis = d
				sum = nums[min] + nums[mid] + nums[max]
			}
			if nums[mid]+nums[max] < t {
				mid++
			} else {
				max--
			}
		}
	}
	return sum
}

// @lc code=end

