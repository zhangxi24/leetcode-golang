/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 *
 * https://leetcode.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (46.74%)
 * Likes:    1454
 * Dislikes: 129
 * Total Accepted:    365.9K
 * Total Submissions: 782.7K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a linked list, swap every two adjacent nodes and return its head.
 *
 * You may not modify the values in the list's nodes, only nodes itself may be
 * changed.
 *
 *
 *
 * Example:
 *
 *
 * Given 1->2->3->4, you should return the list as 2->1->4->3.
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	ret := &ListNode{Val: 0, Next: nil}
	h := ret
	pre := head
	for pre != nil && pre.Next != nil {
		now := pre.Next
		ret.Next = now
		ret = ret.Next
		now = now.Next

		ret.Next = pre
		ret = ret.Next
		ret.Next = nil
		pre = now
	}
	if pre != nil {
		ret.Next = pre
	}

	return h.Next
}

// @lc code=end

