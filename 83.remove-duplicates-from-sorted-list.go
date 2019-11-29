/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (43.55%)
 * Likes:    984
 * Dislikes: 88
 * Total Accepted:    381.3K
 * Total Submissions: 872.2K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a sorted linked list, delete all duplicates such that each element
 * appear only once.
 *
 * Example 1:
 *
 *
 * Input: 1->1->2
 * Output: 1->2
 *
 *
 * Example 2:
 *
 *
 * Input: 1->1->2->3->3
 * Output: 1->2->3
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
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	h := &ListNode{}
	cur := h
	fast := head
	for fast != nil {
		tmp := fast
		fast = fast.Next
		for fast != nil && fast.Val == tmp.Val {
			fast = fast.Next
		}
		cur.Next = tmp
		cur = cur.Next
		cur.Next = nil
	}
	return h.Next
}

// @lc code=end

