package medium

/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	for _, v := range prerequisites {
		indegree[v[0]]++
	}
	degree := []int{}
	for i, v := range indegree {
		if v == 0 {
			degree = append(degree, i)
		}
	}
	count := len(degree)
	for len(degree) > 0 {
		cur := degree[0]
		degree = degree[1:]

		for _, v := range prerequisites {
			if v[1] == cur {
				indegree[v[0]]--
				if indegree[v[0]] == 0 {
					count++
					degree = append(degree, v[0])
				}
			}
		}
	}
	return count == numCourses
}

// @lc code=end
