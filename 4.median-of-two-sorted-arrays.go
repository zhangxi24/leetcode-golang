/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (27.53%)
 * Likes:    5260
 * Dislikes: 768
 * Total Accepted:    530.9K
 * Total Submissions: 1.9M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 *
 * Find the median of the two sorted arrays. The overall run time complexity
 * should be O(log (m+n)).
 *
 * You may assume nums1 and nums2 cannot be both empty.
 *
 * Example 1:
 *
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * The median is 2.0
 *
 *
 * Example 2:
 *
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * The median is (2 + 3)/2 = 2.5
 *
 *
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var mid float64
	mini := 0
	minj := 0
	maxi := len(nums1) - 1
	maxj := len(nums2) - 1
	avg := (len(nums1) + len(nums2)) % 2
	for mini <= maxi && minj <= maxj {
		// 只有两个数了
		if mini == maxi && minj == maxj {
			break
		}
		if nums1[mini] < nums2[minj] {
			mini++
		} else {
			minj++
		}
		if nums1[maxi] < nums2[maxj] {
			maxj--
		} else {
			maxi--
		}
	}
	if mini == maxi && minj == maxj {
		mid = (float64(nums1[mini]) + float64(nums2[minj])) / 2
	} else if mini > maxi { // 只有数组2有数
		if avg == 0 { // 需要求均值
			mid = (float64(nums2[(minj+maxj+1)/2]) + float64(nums2[(minj+maxj-1)/2])) / 2
		} else { // 直接是中位数
			mid = float64(nums2[(minj+maxj)/2])
		}
	} else { // 只有数组1有数
		if avg == 0 {
			mid = (float64(nums1[(mini+maxi+1)/2]) + float64(nums1[(mini+maxi-1)/2])) / 2
		} else {
			mid = float64(nums1[(mini+maxi)/2])
		}
	}
	return mid
}

// @lc code=end

