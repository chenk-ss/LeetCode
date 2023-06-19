package medium

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func letterCombinations(digits string) []string {
	m := map[int]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}
	result := []string{}
	for _, v := range digits {
		strs := m[int(v-'0')]
		mid := []string{}
		for _, sv := range strs {
			if len(result) == 0 {
				mid = append(mid, string(sv))
			} else {
				for _, mv := range result {
					mid = append(mid, mv+string(sv))
				}
			}
		}
		result = mid[:]
	}
	return result
}

// @lc code=end
