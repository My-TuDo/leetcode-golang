/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	// 1.初始化
	indegree := make([]int, numCourses)
	adj := make([][]int, numCourses)
	for _, pre := range prerequisites {
		cur, next := pre[0], pre[1]
		indegree[cur]++
		adj[next] = append(adj[next], cur)
	}

	// 2. 找到说有入度为0的课程
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 3. 拓扑排序
	res := make([]int, 0)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		res = append(res, cur)

		for _, next := range adj[cur] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	if len(res) != numCourses {
		return []int{}
	}
	return res
}

// @lc code=end

