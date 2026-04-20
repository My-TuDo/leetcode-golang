/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 1. 使用快慢指针找到中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. 切断链表，并反转后半部分
	mid := slow.Next
	slow.Next = nil // 断开前半部分和后半部分
	newBack := reverse(mid)

	// 3. 交替合并两个链表
	curr1 := head
	curr2 := newBack
	for curr2 != nil {
		// 暂存下一个节点，防止断链
		next1 := curr1.Next
		next2 := curr2.Next

		curr1.Next = curr2
		curr2.Next = next1

		// 移动指针
		curr1 = next1
		curr2 = next2
	}
}

// 辅助函数：反转链表
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// @lc code=end

