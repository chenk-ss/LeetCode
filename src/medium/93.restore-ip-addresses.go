package medium

import "strings"

/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	result := []string{}
	restoreIpAddressesFunc(s, []string{}, &result)
	return result
}

func restoreIpAddressesFunc(s string, list []string, res *[]string) {
	if len(list) == 4 && len(s) > 0 {
		return
	}
	if len(list) < 4 && len(s) == 0 {
		return
	}
	if len(s) == 0 && len(list) == 4 {
		*res = append(*res, strings.Join(list, "."))
		return
	}
	restoreIpAddressesFunc(s[1:], append(list, string(s[0])), res)
	if s[0] != '0' {
		if len(s) > 1 {
			restoreIpAddressesFunc(s[2:], append(list, s[:2]), res)
		}
		if len(s) > 2 && s[:3] >= "100" && s[:3] <= "255" {
			restoreIpAddressesFunc(s[3:], append(list, s[:3]), res)
		}
	}
}

// @lc code=end
