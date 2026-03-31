/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head} // 创建一个虚拟头节点，指向原链表的头节点
	current := dummy               // 从虚拟头节点开始遍历链表
	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next // 跳过当前节点，删除它
		} else {
			current = current.Next // 继续遍历下一个节点
		}
	}
	return dummy.Next // 返回新的链表头节点
}

// @lc code=end

