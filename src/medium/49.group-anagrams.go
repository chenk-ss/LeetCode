package medium

import "strconv"

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		str := groupAnagramsFunc(v)
		if _, ok := m[str]; ok {
			m[str] = append(m[str], v)
		} else {
			m[str] = []string{v}
		}
	}

	result := [][]string{}
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func groupAnagramsFunc(str string) string {
	result := ""
	list := make([]int, 26)
	for _, v := range str {
		list[int(byte(v)-'a')]++
	}
	for i, v := range list {
		if v > 0 {
			result += strconv.Itoa(v) + string(i+'a')
		}
	}
	return result
}

// @lc code=end
