/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	// 堆排序
	heap := nums[:k]
	buildMinHeap(heap)

	for i := k; i < len(nums); i++ {
		// 如果当前元素比堆顶大，说明它可以进入
		if nums[i] > heap[0] {
			heap[0] = nums[i]
			minHeapify(heap, 0, k)
		}
	}
	return heap[0]
}

// 建立初始化小顶堆
func buildMinHeap(a []int) {
	n := len(a)
	for i := n/2 - 1; i >= 0; i-- {
		minHeapify(a, i, n)
	}
}

// 堆化
func minHeapify(a []int, i, n int) {
	left, right := i*2+1, i*2+2
	smallest := i

	if left < n && a[left] < a[smallest] {
		smallest = left
	}
	if right < n && a[right] < a[smallest] {
		smallest = right
	}

	if smallest != i {
		a[i], a[smallest] = a[smallest], a[i]
		minHeapify(a, smallest, n)
	}
}

// @lc code=end

