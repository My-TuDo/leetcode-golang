/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	// 本质是一个链表问题，快慢指针来判断是否有环
	slow, fast := n, n
	for {
		slow = getNext(slow)          // 慢指针走一步
		fast = getNext(getNext(fast)) // 快指针走两步
		if slow == fast {             // 如果相遇，说明有环
			break
		}
	}
	return slow == 1 // 如果相遇点是1，则是快乐数

}
func getNext(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

// @lc code=end

