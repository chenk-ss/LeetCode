package medium

/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	res := []int{}
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
	res = append(res, degree...)
	for len(degree) > 0 {
		cur := degree[0]
		degree = degree[1:]
		for _, v := range prerequisites {
			if v[1] == cur {
				indegree[v[0]]--
				if indegree[v[0]] == 0 {
					res = append(res, v[0])
					degree = append(degree, v[0])
				}
			}
		}
	}
	if len(res) != numCourses {
		return nil
	}
	return res
}

// @lc code=end
