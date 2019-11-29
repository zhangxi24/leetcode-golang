/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
 *
 * https://leetcode.com/problems/recover-binary-search-tree/description/
 *
 * algorithms
 * Hard (35.99%)
 * Likes:    1045
 * Dislikes: 58
 * Total Accepted:    135.3K
 * Total Submissions: 371.1K
 * Testcase Example:  '[1,3,null,null,2]'
 *
 * Two elements of a binary search tree (BST) are swapped by mistake.
 *
 * Recover the tree without changing its structure.
 *
 * Example 1:
 *
 *
 * Input: [1,3,null,null,2]
 *
 * 1
 * /
 * 3
 * \
 * 2
 *
 * Output: [3,1,null,null,2]
 *
 * 3
 * /
 * 1
 * \
 * 2
 *
 *
 * Example 2:
 *
 *
 * Input: [3,1,4,null,null,2]
 *
 * ⁠ 3
 * ⁠/ \
 * 1   4
 * /
 * 2
 *
 * Output: [2,1,4,null,null,3]
 *
 * ⁠ 2
 * ⁠/ \
 * 1   4
 * /
 * ⁠ 3
 *
 *
 * Follow up:
 *
 *
 * A solution using O(n) space is pretty straight forward.
 * Could you devise a constant space solution?
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

func recoverTree(root *TreeNode) {
	var a, b, pre *TreeNode
	helper(root, &pre, &a, &b)
	a.Val, b.Val = b.Val, a.Val
}

func helper(root *TreeNode, pre, a, b **TreeNode) {
	if root == nil {
		return
	}

	helper(root.Left, pre, a, b)
	if *pre != nil && (*pre).Val >= root.Val { // mark
		if *a == nil {
			*a = *pre
			*b = root
		} else {
			*b = root
		}
	}
	*pre = root
	helper(root.Right, pre, a, b)
}

// @lc code=end

