/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	// 使用map存储nums1中的元素
	set := make(map[int]bool)
	for _, num := range nums1 {
		set[num] = true
	}
	// 定义结果切片
	result := []int{}
	// 遍历nums2，检查元素是否在set中，如果存在则添加到结果中，并从set中删除以避免重复
	for _, num := range nums2 {
		if set[num] {
			result = append(result, num)
			delete(set, num)
		}
	}
	return result

}

// @lc code=end

