package medium

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 */

// @lc code=start
func frequencySort(s string) string {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	list := []rune{}
	for k, _ := range m {
		list = append(list, k)
	}
	sort.Slice(list, func(i, j int) bool {
		return m[list[i]] > m[list[j]]
	})
	res := ""
	for _, v := range list {
		res += strings.Repeat(string(v), m[v])
	}
	return res
}

// @lc code=end
