package medium

/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	list1, list2 := []int{}, []int{}
	count := 0
	for i := len(num2) - 1; i >= 0; i-- {
		v := int(num2[i] - '0')

		list2 = []int{}
		for j := 0; j < count; j++ {
			list2 = append(list2, 0)
		}
		count++

		n1 := num1
		extraNum := 0
		for len(n1) > 0 {
			v1 := int(n1[len(n1)-1] - '0')
			newNum := v*v1 + extraNum
			if newNum > 9 {
				extraNum = newNum / 10
				list2 = append(list2, newNum%10)
			} else {
				extraNum = 0
				list2 = append(list2, newNum)
			}
			n1 = n1[:len(n1)-1]
		}
		if extraNum != 0 {
			list2 = append(list2, extraNum)
		}
		list3 := []int{}
		extra := 0
		for len(list1) > 0 || len(list2) > 0 {
			sum := 0
			if len(list1) > 0 {
				sum += list1[0]
				list1 = list1[1:]
			}
			if len(list2) > 0 {
				sum += list2[0]
				list2 = list2[1:]
			}
			sum += extra
			if sum > 9 {
				list3 = append(list3, sum%10)
				extra = sum / 10
			} else {
				list3 = append(list3, sum)
				extra = 0
			}
		}
		if extra != 0 {
			list3 = append(list3, extra)
		}
		list1 = list3
	}
	for i, j := 0, len(list1)-1; i < j; i, j = i+1, j-1 {
		list1[i], list1[j] = list1[j], list1[i]
	}
	result := ""
	for _, v := range list1 {
		result += string(v + '0')
	}
	return result
}

// @lc code=end
