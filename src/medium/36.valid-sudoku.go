package medium

/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		list1 := make([]int, 9)
		for j := 0; j < 9; j++ {
			v := board[i][j]
			if v != '.' {
				if list1[v-'1'] != 0 {
					return false
				} else {
					list1[v-'1'] = 1
				}
			}
		}

		list2 := make([]int, 9)
		for j := 0; j < 9; j++ {
			v := board[j][i]
			if v != '.' {
				if list2[v-'1'] != 0 {
					return false
				} else {
					list2[v-'1'] = 1
				}
			}
		}
	}
	count := 1
	for count <= 9 {
		maxI := (count * 3) % 9
		if count%3 == 0 {
			maxI = 9
		}
		list := make([]int, 9)
		for i := ((count - 1) * 3) % 9; i < maxI; i++ {
			for j := (count - 1) / 3 * 3; j < (count-1)/3*3+3; j++ {
				v := board[i][j]
				if v != '.' {
					if list[v-'1'] != 0 {
						return false
					} else {
						list[v-'1'] = 1
					}
				}
			}
		}
		count++
	}

	return true
}

// @lc code=end
