package medium

/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
type LRUCache struct {
	m    map[int]int
	list []int
	cap  int
}

func Constructor146(capacity int) LRUCache {
	return LRUCache{m: make(map[int]int), list: []int{}, cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.m[key]
	if !ok {
		return -1
	}
	this.ReOrder(key)
	return v
}

func (this *LRUCache) ReOrder(key int) {
	index := -1
	for i := range this.list {
		if this.list[i] == key {
			index = i
			break
		}
	}
	list := []int{}
	list = append(list, this.list[:index]...)
	list = append(list, this.list[index+1:]...)
	list = append(list, key)
	this.list = list
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.m[key]
	this.m[key] = value
	if ok {
		this.ReOrder(key)
	} else if len(this.list) == this.cap {
		delete(this.m, this.list[0])
		this.list = append(append([]int{}, this.list[1:]...), key)
	} else {
		this.list = append(this.list, key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
