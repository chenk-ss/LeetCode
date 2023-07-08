package medium

/*
 * @lc app=leetcode id=433 lang=golang
 *
 * [433] Minimum Genetic Mutation
 */

// @lc code=start
func minMutation(startGene string, endGene string, bank []string) int {
	m := make(map[string]bool)
	for _, v := range bank {
		m[v] = true
	}
	count := 0
	list := []string{startGene}
	for len(list) > 0 {
		tmp := []string{}
		for _, v := range list {
			if v == endGene {
				return count
			}
			for mKey := range m {
				if minMutationFunc(v, mKey) == 1 {
					tmp = append(tmp, mKey)
					delete(m, mKey)
				}
			}
		}
		count++
		list = tmp
	}
	return -1
}

func minMutationFunc(a, b string) int {
	count := 0
	for i := 0; i < 8; i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count
}

// @lc code=end
