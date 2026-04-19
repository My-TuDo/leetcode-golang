/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并 K 个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	// 分治法
	mid := len(lists) / 2

	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])

	return mergeTwoLists(left, right)
}

func mergeTwoLists(L1, L2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for L1 != nil && L2 != nil {
		if L1.Val < L2.Val {
			curr.Next = L1
			L1 = L1.Next
		} else {
			curr.Next = L2
			L2 = L2.Next
		}
		curr = curr.Next
	}
	if L1 != nil {
		curr.Next = L1
	}
	if L2 != nil {
		curr.Next = L2
	}
	return dummy.Next
}

// @lc code=end

