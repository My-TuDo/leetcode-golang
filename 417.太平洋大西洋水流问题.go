/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] 太平洋大西洋水流问题
 */

// @lc code=start
func pacificAtlantic(heights [][]int) [][]int {
	// 逆流，从边缘出发寻找最高点。
	// 分别从上下和左右进行dfs，然后取交集
	if len(heights) == 0 {
		return nil
	}
	m, n := len(heights), len(heights[0])

	// 两个布尔矩阵，记录算法能流向对应的海洋
	canReachP := make([][]bool, m)
	canReachA := make([][]bool, m)
	for i := range canReachP {
		canReachP[i] = make([]bool, n)
		canReachA[i] = make([]bool, n)
	}

	var dfs func(r, c int, reachable [][]bool)
	dfs = func(r, c int, reachable [][]bool) {
		reachable[r][c] = true
		// 四个方向：上、下、左、右
		dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			// 越界检查 是否访问过 逆流条件 新高度 >= 旧高度
			if nr >= 0 && nr < m && nc >= 0 && nc < n && !reachable[nr][nc] && heights[nr][nc] >= heights[r][c] {
				dfs(nr, nc, reachable)
			}
		}
	}
	// 1. 从左右边开始dfs
	for i := 0; i < m; i++ {
		dfs(i, 0, canReachP)   // 左边：太平洋
		dfs(i, n-1, canReachA) // 右边：大西洋
	}
	// 2. 从上下边界开始dfs
	for j := 0; j < n; j++ {
		dfs(0, j, canReachP)   // 上边： 太平洋
		dfs(m-1, j, canReachA) // 下边： 大西洋
	}

	// 3. 找出交集
	res := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if canReachP[i][j] && canReachA[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

// @lc code=end

