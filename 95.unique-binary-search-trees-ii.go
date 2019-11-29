/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 *
 * https://leetcode.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (37.26%)
 * Likes:    1609
 * Dislikes: 133
 * Total Accepted:    160.3K
 * Total Submissions: 426K
 * Testcase Example:  '3'
 *
 * Given an integer n, generate all structurally unique BST's (binary search
 * trees) that store values 1 ... n.
 *
 * Example:
 *
 *
 * Input: 3
 * Output:
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * Explanation:
 * The above output corresponds to the 5 unique BST's shown below:
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return helper(1, n)
}

func helper(i, n int) []*TreeNode {
	if i > n {
		return []*TreeNode{nil}
	}

	trees := []*TreeNode{}
	for j := i; j <= n; j++ {
		ltrees := helper(i, j-1)
		rtrees := helper(j+1, n)
		for _, l := range ltrees {
			for _, r := range rtrees {
				trees = append(trees, &TreeNode{Val: j, Left: l, Right: r})
			}
		}
	}
	return trees
}

// @lc code=end

