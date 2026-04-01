/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 创建一个虚拟头节点，方便操作
	dummy := &ListNode{}
	current := dummy
	// 同时遍历两个链表，比较当前节点的值
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1 // 将较小的节点连接到结果链表
			list1 = list1.Next   // 移动 list1 的指针
		} else {
			current.Next = list2 // 将较小的节点连接到结果链表
			list2 = list2.Next   // 移动 list2 的指针
		}
		current = current.Next // 移动结果链表的指针
	}
	// 如果其中一个链表还有剩余节点，直接连接到结果链表末尾
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}
	return dummy.Next // 返回合并后的链表头节点
}

// @lc code=end

