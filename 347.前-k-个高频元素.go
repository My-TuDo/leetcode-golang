/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	// map计数 + 桶排序
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	// 桶：下标是频率，值是具有该频率的数字切片
	buckets := make([][]int, len(nums)+1)
	for num, freq := range countMap {
		buckets[freq] = append(buckets[freq], num)
	}

	res := make([]int, 0, k)
	// 从大到小遍历桶
	for i := len(buckets) - 1; i >= 0 && len(res) < k; i-- {
		if len(buckets[i]) > 0 {
			res = append(res, buckets[i]...) // 将一个切片中的元素一个一个打包进另一个切片中。
		}
	}
	return res
}

// @lc code=end

