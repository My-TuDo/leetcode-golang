/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make([]int, numCourses)
	adjacency := make([][]int, numCourses)

	// 1. 建立邻接表并统计入度
	for _, pre := range prerequisites {
		course, preCourse := pre[0], pre[1]
		inDegree[course]++
		adjacency[preCourse] = append(adjacency[preCourse], course)
	}

	// 2. 将所有入度为0的节点入队
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 3. BFS拓扑排序
	count := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		count++

		// 遍历当前课程的所有后续课程
		for _, nextCourse := range adjacency[curr] {
			inDegree[nextCourse]--
			if inDegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}

	// 4. 如果修过的课等于总数，说明没环
	return count == numCourses
}

// @lc code=end

