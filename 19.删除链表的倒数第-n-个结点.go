/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
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
	dummy := &ListNode{Next: head} // 创建一个虚拟头节点，指向原链表的头节点
	first, second := dummy, dummy  // 定义两个指针，初始都指向虚拟头节点
	// 先让 first 指针向前移动 n 步
	for i := 0; i < n; i++ {
		first = first.Next
	}
	// 然后同时移动 first 和 second 指针，直到 first 指针到达链表末尾
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	// 此时 second 指针指向要删除节点的前一个节点，跳过要删除的节点
	second.Next = second.Next.Next
	return dummy.Next // 返回新的链表头节点
}

// @lc code=end

