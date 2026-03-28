/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	for _, v := range nums {
		index := abs(v)

		nums[index-1] = -abs(nums[index-1])
	}
	var reslt []int
	for i, v := range nums {
		if v > 0 {
			reslt = append(reslt, i+1)
		}
	}
	return reslt
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

