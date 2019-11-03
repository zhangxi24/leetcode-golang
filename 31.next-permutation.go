import "fmt"

/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 *
 * https://leetcode.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (31.19%)
 * Likes:    2310
 * Dislikes: 754
 * Total Accepted:    283.1K
 * Total Submissions: 906.6K
 * Testcase Example:  '[1,2,3]'
 *
 * Implement next permutation, which rearranges numbers into the
 * lexicographically next greater permutation of numbers.
 *
 * If such arrangement is not possible, it must rearrange it as the lowest
 * possible order (ie, sorted in ascending order).
 *
 * The replacement must be in-place and use only constant extra memory.
 *
 * Here are some examples. Inputs are in the left-hand column and its
 * corresponding outputs are in the right-hand column.
 *
 * 1,2,3 → 1,3,2
 * 3,2,1 → 1,2,3
 * 1,1,5 → 1,5,1
 *
 */

// @lc code=start
func nextPermutation(nums []int) {
	dst := len(nums) - 1
	src := dst
	find := false
	for dst > 0 {
		if nums[dst] <= nums[dst-1] { // 找到需要更换数字的最高位
			dst--
			continue
		}
		dst--
		for ; nums[src] <= nums[dst]; src-- { // 找到应放在最高位的数字当前所在位置
		}
		find = true
		break
	}

	i := 0
	if find == true { // src与dst位置交换，之后逆序[dst+1,len(nums)-1]
		nums[dst], nums[src] = nums[src], nums[dst]
		i = dst + 1
	} // 否则整串逆序
	posSum := i + len(nums) - 1
	for ; i <= posSum/2; i++ {
		fmt.Println(i, posSum-i)
		nums[i], nums[posSum-i] = nums[posSum-i], nums[i]
	}
}

// @lc code=end

