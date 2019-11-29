/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 *
 * https://leetcode.com/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (48.14%)
 * Likes:    2287
 * Dislikes: 87
 * Total Accepted:    234.5K
 * Total Submissions: 481.9K
 * Testcase Example:  '3'
 *
 * Given n, how many structurally unique BST's (binary search trees) that store
 * values 1 ... n?
 *
 * Example:
 *
 *
 * Input: 3
 * Output: 5
 * Explanation:
 * Given n = 3, there are a total of 5 unique BST's:
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 *
 *
 */

// @lc code=start
func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	num := make([]int, n+1)
	num[0] = 1
	num[1] = 1
	for i := 2; i <= n; i++ { // root
		for j := 1; j <= i; j++ { // j is subtree's root
			num[i] += num[j-1] * num[i-j]
		}
	}
	return num[n]
}

// @lc code=end

