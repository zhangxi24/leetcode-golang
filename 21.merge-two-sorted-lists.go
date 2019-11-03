/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (49.42%)
 * Likes:    2801
 * Dislikes: 410
 * Total Accepted:    724.5K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	l := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			l.Next = l1
			l1 = l1.Next
			l = l.Next
		} else {
			l.Next = l2
			l2 = l2.Next
			l = l.Next
		}
	}
	if l1 != nil {
		l.Next = l1
	}
	if l2 != nil {
		l.Next = l2
	}
	return head.Next
}

// @lc code=end

