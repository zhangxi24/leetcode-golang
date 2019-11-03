/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (34.62%)
 * Likes:    2255
 * Dislikes: 167
 * Total Accepted:    470.3K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, remove the n-th node from the end of list and return
 * its head.
 *
 * Example:
 *
 *
 * Given linked list: 1->2->3->4->5, and n = 2.
 *
 * After removing the second node from the end, the linked list becomes
 * 1->2->3->5.
 *
 *
 * Note:
 *
 * Given n will always be valid.
 *
 * Follow up:
 *
 * Could you do this in one pass?
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first, pre, del, h := head, head, head, head
	for i := 0; i < n; i++ {
		pre = pre.Next
	}
	j := 0
	for pre != nil {
		pre = pre.Next
		del = del.Next
		if j != 0 {
			h = h.Next
		}
		j++
	}
	if del == first {
		first = first.Next
	} else {
		h.Next = del.Next
	}

	return first
}

// @lc code=end

