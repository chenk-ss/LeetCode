package medium

/*
 * @lc app=leetcode id=134 lang=golang
 *
 * [134] Gas Station
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	gasLeft, fullGasLeft, startIndex := 0, 0, 0
	for i := 0; i < length; i++ {
		gasLeft += gas[i] - cost[i]
		fullGasLeft += gas[i] - cost[i]
		if gasLeft < 0 {
			startIndex = i + 1
			gasLeft = 0
		}
	}
	if fullGasLeft < 0 {
		return -1
	}
	return startIndex
}

// @lc code=end
