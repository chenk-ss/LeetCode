package medium

/*
 * @lc app=leetcode id=399 lang=golang
 *
 * [399] Evaluate Division
 */

// @lc code=start
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// https://leetcode.com/problems/evaluate-division/solutions/2489823/go-0-ms-faster-than-100/
	g := make(map[string]map[string]float64)
	for i, v := range equations {
		if g[v[0]] == nil {
			g[v[0]] = make(map[string]float64)
		}
		if g[v[1]] == nil {
			g[v[1]] = make(map[string]float64)
		}
		g[v[0]][v[1]] = values[i]
		g[v[1]][v[0]] = 1 / values[i]
	}

	results := make([]float64, 0, len(queries))
	for _, v := range queries {
		x := v[0]
		y := v[1]
		result := calcEquationFunc(x, y, make(map[string]bool), g)
		results = append(results, result)
	}
	return results
}

func calcEquationFunc(x, y string, visited map[string]bool, g map[string]map[string]float64) float64 {
	_, okX := g[x]
	_, okY := g[y]
	if !okX || !okY {
		return -1.0
	}
	if len(g[x]) == 0 {
		return -1.0
	}
	val, okResult := g[x][y]
	if okResult {
		return val
	}
	for k, _ := range g[x] {
		if !visited[k] {
			visited[k] = true
			tmp := calcEquationFunc(k, y, visited, g)
			if tmp == -1.0 {
				continue
			} else {
				return tmp * g[x][k]
			}
		}
	}
	return -1.0
}

// @lc code=end
