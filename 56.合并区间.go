/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 按照左边界升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]

		// 如果当前区间的左边界 <= 前一个区间的右边界，说明重叠
		if curr[0] <= prev[1] {
			// 合并
			if curr[1] > prev[1] {
				prev[1] = curr[1]
			}
		} else {
			// 不重叠，把前一个区间存入结果，更新prev
			res = append(res, prev)
			prev = curr
		}
	}
	res = append(res, prev)
	return res
}

// @lc code=end

