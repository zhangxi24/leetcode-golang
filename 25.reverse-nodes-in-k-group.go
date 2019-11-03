/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 *
 * https://leetcode.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (38.21%)
 * Likes:    1462
 * Dislikes: 295
 * Total Accepted:    213.5K
 * Total Submissions: 557K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, reverse the nodes of a linked list k at a time and
 * return its modified list.
 *
 * k is a positive integer and is less than or equal to the length of the
 * linked list. If the number of nodes is not a multiple of k then left-out
 * nodes in the end should remain as it is.
 *
 *
 *
 *
 * Example:
 *
 * Given this linked list: 1->2->3->4->5
 *
 * For k = 2, you should return: 2->1->4->3->5
 *
 * For k = 3, you should return: 3->2->1->4->5
 *
 * Note:
 *
 *
 * Only constant extra memory is allowed.
 * You may not alter the values in the list's nodes, only nodes itself may be
 * changed.
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
func reverse(ret *ListNode, l []*ListNode, k int) *ListNode {
	for i := 0; i < k; i++ {
		ret.Next = l[k-i-1]
		ret = ret.Next
	}
	ret.Next = nil
	return ret
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	ret := &ListNode{Val: 0, Next: nil}
	h := ret
	pre := head
	l := make([]*ListNode, k)
	for pre != nil {
		i := 0
		n := pre
		for ; i < k; i++ {
			if pre != nil {
				l[i] = pre
				pre = pre.Next
			} else {
				break
			}
		}
		if i != k {
			ret.Next = n
		} else {
			ret = reverse(ret, l, k)
		}
	}

	return h.Next
}

// @lc code=end

