package medium

import "strings"

/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
func simplifyPath(path string) string {
	result := []string{}
	for _, v := range strings.Split(path, "/") {
		if v == "" || v == "." {
			continue
		} else if v != ".." {
			result = append(result, v)
		} else if len(result) > 0 {
			result = result[:len(result)-1]
		}
	}
	return "/" + strings.Join(result, "/")
}

// @lc code=end
