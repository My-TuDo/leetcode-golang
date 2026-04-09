/*
 * @lc app=leetcode.cn id=733 lang=golang
 *
 * [733] 图像渲染
 */

// @lc code=start
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == color {
		return image
	}

	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) || image[r][c] != oldColor {
			return
		}
		image[r][c] = color
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	dfs(sr, sc)
	return image
}

// @lc code=end

