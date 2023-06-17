package easy

/*
 * @lc app=leetcode id=495 lang=golang
 *
 * [495] Teemo Attacking
 */

// @lc code=start
func findPoisonedDuration(timeSeries []int, duration int) int {
	sum := duration
	for i := range timeSeries {
		if i != len(timeSeries)-1 {
			if timeSeries[i]+duration <= timeSeries[i+1] {
				sum += duration
			} else {
				sum += timeSeries[i+1] - timeSeries[i]
			}
		}
	}
	return sum
}

// @lc code=end
