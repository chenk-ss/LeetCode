package medium

import (
	"sort"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 */

// @lc code=start
func largestNumber(nums []int) string {
	list := []string{}
	for _, v := range nums {
		list = append(list, strconv.Itoa(v))
	}
	sort.Slice(list, func(a, b int) bool {
		return list[a]+list[b] > list[b]+list[a]
	})
	if list[0] == "0" {
		return "0"
	}
	return strings.Join(list, "")
}

// @lc code=end
