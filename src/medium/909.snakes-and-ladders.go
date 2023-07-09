package medium

/*
 * @lc app=leetcode id=909 lang=golang
 *
 * [909] Snakes and Ladders
 */

// @lc code=start
func snakesAndLadders(board [][]int) int {
	n := len(board)
	steps := make([]int, n*n)

	var snakesAndLaddersDfs func(from int)
	snakesAndLaddersDfs = func(from int) {
		step := steps[from]
		nextpos := []int{}
		for pos := from + 1; pos <= from+6; pos++ {
			if pos >= n*n {
				break
			}
			x, y := snakesAndLaddersFunc(n, pos)
			if board[x][y] != -1 {
				topos := board[x][y] - 1
				if steps[topos] == 0 || step+1 < steps[topos] {
					steps[topos] = step + 1
					nextpos = append(nextpos, topos)
				}
			} else if steps[pos] == 0 || step+1 < steps[pos] {
				steps[pos] = step + 1
				nextpos = append(nextpos, pos)
			}
		}
		for _, pos := range nextpos {
			snakesAndLaddersDfs(pos)
		}
	}
	snakesAndLaddersDfs(0)
	if step := steps[n*n-1]; step != 0 {
		return step
	}
	return -1
}

func snakesAndLaddersFunc(n, num int) (int, int) {
	a := num / n
	b := num % n
	if a%2 == 1 {
		b = n - b - 1
	}
	return n - a - 1, b
}

// @lc code=end
