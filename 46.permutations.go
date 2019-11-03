/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 *
 * https://leetcode.com/problems/permutations/description/
 *
 * algorithms
 * Medium (57.98%)
 * Likes:    2614
 * Dislikes: 84
 * Total Accepted:    459.7K
 * Total Submissions: 788.5K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a collection of distinct integers, return all possible permutations.
 *
 * Example:
 *
 *
 * Input: [1,2,3]
 * Output:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 *
 */

// @lc code=start

// 回溯
func permute(nums []int) [][]int {
	res := [][]int{}
	used := map[int]bool{}
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
			used[j] = true
			r = append(r, n)
			helper(nums, depth, i+1, used, r, res)
			r = r[:len(r)-1]
			used[j] = false
		}
	}
}

/*
func permute(nums []int) [][]int {
	res := [][]int{}
	helper(&res, []int{}, nums)
	return res
}

func helper(res *[][]int, r, num []int) {
	if len(num) == 0 {
		tempp := make([]int, len(r))
		copy(tempp, r)
		*res = append(*res, tempp)
		return
	}

	for i := 0; i < len(num); i++ {
		r1 := append(r, num[i])
		ans := append([]int{}, num[:i]...)
		ans = append(ans, num[i+1:]...)
		helper(res, r1, ans)
	}

}
*/

// @lc code=end

