package easy

import "strconv"

/*
 * @lc app=leetcode id=401 lang=golang
 *
 * [401] Binary Watch
 */

// @lc code=start
func readBinaryWatch(turnedOn int) []string {
	result := []string{}
	list := make([]int, 60)
	for i := 1; i < 60; i++ {
		list[i] = list[i/2] + i&1
	}
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if list[i]+list[j] == turnedOn {
				if j < 10 {
					result = append(result, strconv.Itoa(i)+":0"+strconv.Itoa(j))
				} else {
					result = append(result, strconv.Itoa(i)+":"+strconv.Itoa(j))
				}
			}
		}
	}
	return result
}

// @lc code=end
