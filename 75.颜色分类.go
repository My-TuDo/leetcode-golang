/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	p0 := 0
	p2 := len(nums) - 1
	curr := 0
	for curr <= p2 {
		if nums[curr] == 0 {
			nums[p0], nums[curr] = nums[curr], nums[p0]
			p0++
			curr++
		} else if nums[curr] == 2 {
			nums[p2], nums[curr] = nums[curr], nums[p2]
			p2--
		} else {
			curr++
		}
	}
}

// @lc code=end

