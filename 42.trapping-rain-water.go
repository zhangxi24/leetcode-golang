/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 *
 * https://leetcode.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (45.20%)
 * Likes:    4773
 * Dislikes: 83
 * Total Accepted:    372.3K
 * Total Submissions: 819.8K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * Given n non-negative integers representing an elevation map where the width
 * of each bar is 1, compute how much water it is able to trap after raining.
 *
 *
 * The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
 * In this case, 6 units of rain water (blue section) are being trapped. Thanks
 * Marcos for contributing this image!
 *
 * Example:
 *
 *
 * Input: [0,1,0,2,1,0,1,3,2,1,2,1]
 * Output: 6
 *
 */

// @lc code=start
func trap(height []int) int {
	area := 0

	// 先找到最高的位置
	pos, h := 0, 0
	for i := 0; i < len(height); i++ {
		if height[i] > h {
			pos = i
			h = height[i]
		}
	}
	// 分别求两侧的面积，边遍历边累积求面积
	lMax := 0
	for i := 0; i < pos; i++ {
		if lMax < height[i] {
			lMax = height[i]
		}
		area += lMax - height[i]
	}
	rMax := 0
	for i := len(height) - 1; i > pos; i-- {
		if rMax < height[i] {
			rMax = height[i]
		}
		area += rMax - height[i]
	}

	return area
}

// @lc code=end

