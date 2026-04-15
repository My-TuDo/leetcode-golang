/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	n := len(nums)

	total := 1 << n
	res := make([][]int, 0, total)

	for i := 0; i < total; i++ {
		path := []int{}
		// 检查数字i的每一位是否为1
		for j := 0; j < n; j++ {
			// 判断i的第j位是否为1
			if (i>>j)&1 == 1 {
				path = append(path, nums[j])
			}
		}
		res = append(res, path)
	}
	return res
}

// @lc code=end

