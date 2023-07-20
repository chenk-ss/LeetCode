package medium

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=165 lang=golang
 *
 * [165] Compare Version Numbers
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	lista := strings.Split(version1, ".")
	listb := strings.Split(version2, ".")
	for i := 0; i < len(lista) || i < len(listb); i++ {
		a, b := 0, 0
		if i < len(lista) {
			a, _ = strconv.Atoi(lista[i])
		}
		if i < len(listb) {
			b, _ = strconv.Atoi(listb[i])
		}
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
	}
	return 0
}

// @lc code=end
