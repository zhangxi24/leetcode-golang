import "sort"

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 *
 * https://leetcode.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (42.62%)
 * Likes:    1335
 * Dislikes: 48
 * Total Accepted:    284.9K
 * Total Submissions: 665K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a collection of numbers that might contain duplicates, return all
 * possible unique permutations.
 *
 * Example:
 *
 *
 * Input: [1,1,2]
 * Output:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 *
 *
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	used := map[int]bool{}
	sort.Ints(nums)
	helper(nums, len(nums), 0, used, []int{}, &res)
	return res
}

func helper(nums []int, depth, i int, used map[int]bool, r []int, res *[][]int) {
	if i == depth {
		tmp := make([]int, len(r))
		copy(tmp, r)
		*res = append(*res, tmp)
		return
	}

	for j, n := range nums {
		if used[j] == false {
			if j > 0 && n == nums[j-1] && used[j-1] {
				continue
			}
			used[j] = true
			r = append(r, n)
			helper(nums, depth, i+1, used, r, res)
			r = r[:len(r)-1]
			used[j] = false
		}
	}
}

// @lc code=end

