/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	// 极速失败:若赎金信长度大于杂志长度，则不可能构成
	if len(ransomNote) > len(magazine) {
		return false
	}

	// 统计杂志中每个字符出现的次数
	charCount := [26]int{}

	// 遍历杂志字符串，统计字符出现次数
	for i := 0; i < len(magazine); i++ {
		charCount[magazine[i]-'a']++
	}

	// 遍历赎金信字符串，检查每个字符是否在杂志中出现足够的次数
	for i := 0; i < len(ransomNote); i++ {
		charCount[ransomNote[i]-'a']--
		if charCount[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}

// @lc code=end

