/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	// 定义快慢指针，初始时都指向链表头节点
	slow, fast := head, head.Next
	// 快慢指针同时移动，slow每次移动一步，fast每次移动两步，直到快指针到达链表末尾或两指针相遇
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

// @lc code=end

