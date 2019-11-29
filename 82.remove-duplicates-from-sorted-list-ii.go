/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (34.34%)
 * Likes:    1115
 * Dislikes: 90
 * Total Accepted:    209.5K
 * Total Submissions: 605.5K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * Given a sorted linked list, delete all nodes that have duplicate numbers,
 * leaving only distinct numbers from the original list.
 *
 * Example 1:
 *
 *
 * Input: 1->2->3->3->4->4->5
 * Output: 1->2->5
 *
 *
 * Example 2:
 *
 *
 * Input: 1->1->1->2->3
 * Output: 2->3
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
	add := h
	ptr := head
	for ptr != nil {
		tmp := ptr
		ptr = ptr.Next
		for ptr != nil && ptr.Val == tmp.Val {
			ptr = ptr.Next
		}
		if ptr == tmp.Next {
			add.Next = tmp
			add = add.Next
			add.Next = nil
		}
	}
	return h.Next
}

// @lc code=end

