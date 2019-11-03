/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 *
 * https://leetcode.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (36.50%)
 * Likes:    3075
 * Dislikes: 205
 * Total Accepted:    470.8K
 * Total Submissions: 1.3M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * Merge k sorted linked lists and return it as one sorted list. Analyze and
 * describe its complexity.
 *
 * Example:
 *
 *
 * Input:
 * [
 * 1->4->5,
 * 1->3->4,
 * 2->6
 * ]
 * Output: 1->1->2->3->4->4->5->6
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
func merge2Lists(a, b *ListNode) *ListNode {
	ret := &ListNode{Val: 0, Next: nil}
	head := ret
	for a != nil && b != nil {
		if a.Val <= b.Val {
			ret.Next = a
			a = a.Next
		} else {
			ret.Next = b
			b = b.Next
		}
		ret = ret.Next
	}
	if a != nil {
		ret.Next = a
	}
	if b != nil {
		ret.Next = b
	}
	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	ret := lists[0]
	for i := 1; i < len(lists); i++ {
		ret = merge2Lists(ret, lists[i])
	}
	return ret
}

// @lc code=end

