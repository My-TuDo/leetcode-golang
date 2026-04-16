/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 */

// @lc code=start
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	count := 0

	// 遍历每一个城市
	for i := 0; i < n; i++ {
		// 如果这个城市还没被归入某个省份
		if !visited[i] {
			// 从这个城市开始深度优先搜索，标记所有相连的城市
			dfs(isConnected, visited, i)
			// 搜索完一圈，省份+1
			count++
		}
	}
	return count
}

func dfs(isConnected [][]int, visited []bool, i int) {
	// 标记当前城市已访问
	visited[i] = true
	// 寻找与城市 i 相连的其他城市 j
	for j := 0; j < len(isConnected); j++ {
		if isConnected[i][j] == 1 && !visited[j] {
			dfs(isConnected, visited, j)
		}
	}
}

// @lc code=end

