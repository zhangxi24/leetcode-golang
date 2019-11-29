/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 *
 * https://leetcode.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (36.29%)
 * Likes:    1603
 * Dislikes: 110
 * Total Accepted:    226.6K
 * Total Submissions: 619.5K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * Reverse a linked list from position m to n. Do it in one-pass.
 *
 * Note: 1 ≤ m ≤ n ≤ length of list.
 *
 * Example:
 *
 *
 * Input: 1->2->3->4->5->NULL, m = 2, n = 4
 * Output: 1->4->3->2->5->NULL
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
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n || head == nil {
		return head
	}

	h := &ListNode{}
	h.Next = head
	pre := h
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	if cur.Next == nil {
		return head
	}

	tail := cur
	next := cur.Next
	for i := m; i < n; i++ {
		tmp := next.Next
		next.Next = cur
		cur = next
		next = tmp
	}
	tail.Next = next
	pre.Next = cur

	return h.Next
}

// @lc code=end

