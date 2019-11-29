/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (58.94%)
 * Likes:    2155
 * Dislikes: 94
 * Total Accepted:    571.1K
 * Total Submissions: 959.5K
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the inorder traversal of its nodes' values.
 *
 * Example:
 *
 *
 * Input: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * Output: [1,3,2]
 *
 * Follow up: Recursive solution is trivial, could you do it iteratively?
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

// 递归
// func inorderTraversal(root *TreeNode) []int {
// 	res := []int{}
// 	if root == nil {
// 		return []int{}
// 	}
// 	res = append(res, inorderTraversal(root.Left)...)
// 	res = append(res, root.Val)
// 	res = append(res, inorderTraversal(root.Right)...)
// 	return res
// }

// 迭代
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	tmp := []*TreeNode{}
	for root != nil || len(tmp) != 0 {
		if root != nil {
			tmp = append(tmp, root)
			root = root.Left
		} else {
			root = tmp[len(tmp)-1]
			res = append(res, root.Val)
			root = root.Right
			tmp = tmp[:len(tmp)-1]
		}
	}

	return res
}

// @lc code=end

