/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 *
 * https://leetcode.com/problems/partition-list/description/
 *
 * algorithms
 * Medium (38.70%)
 * Likes:    879
 * Dislikes: 234
 * Total Accepted:    185K
 * Total Submissions: 474.8K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * Given a linked list and a value x, partition it such that all nodes less
 * than x come before nodes greater than or equal to x.
 *
 * You should preserve the original relative order of the nodes in each of the
 * two partitions.
 *
 * Example:
 *
 *
 * Input: head = 1->4->3->2->5->2, x = 3
 * Output: 1->2->2->4->3->5
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
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	less := &ListNode{}
	more := &ListNode{}
	l, m := less, more
	for h := head; h != nil; h = h.Next {
		tmp := h
		if h.Val < x {
			l.Next = tmp
			l = l.Next
		} else {
			m.Next = tmp
			m = m.Next
		}
	}
	m.Next = nil
	l.Next = more.Next
	return less.Next
}

// @lc code=end

