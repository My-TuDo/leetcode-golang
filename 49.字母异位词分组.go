/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// 初始化一个哈希表
	type anagram [26]int8
	m := make(map[anagram][]string)
	// 遍历每一个字符串元素
	for _, ch := range strs {
		num := anagram{} // 定义一个计数器
		for _, c := range ch {
			num[int(c-'a')]++
		}
		m[num] = append(m[num], ch)
	}

	// 遍历哈希表，存入二维切片中
	str := make([][]string, 0, len(m))
	for _, value := range m {
		str = append(str, value)
	}
	return str
}

// @lc code=end

