package medium

import (
	"math/rand"
)

/*
 * @lc app=leetcode id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */

// @lc code=start
type RandomizedSet struct {
	m    map[int]bool
	vals []int
}

func Constructor380() RandomizedSet {
	return RandomizedSet{m: make(map[int]bool), vals: []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	for _, ok := this.m[val]; ok; {
		return false
	}
	this.m[val] = true
	this.vals = append(this.vals, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	for _, ok := this.m[val]; !ok; {
		return false
	}
	delete(this.m, val)
	index := -1
	for i, v := range this.vals {
		if v == val {
			index = i
		}
	}
	if index == -1 {
		return false
	}
	this.vals = append(this.vals[:index], this.vals[index+1:]...)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.vals[rand.Intn(len(this.vals))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
