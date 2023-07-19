package medium

/*
 * @lc app=leetcode id=211 lang=golang
 *
 * [211] Design Add and Search Words Data Structure
 */

// @lc code=start
type WordDictionary struct {
	children [26]*WordDictionary
	end      bool
}

func Constructor211() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this
	for _, v := range word {
		index := v - 'a'
		if cur.children[index] == nil {
			cur.children[index] = &WordDictionary{}
		}
		cur = cur.children[index]
	}
	cur.end = true
}

func (this *WordDictionary) Search(word string) bool {
	return SearchFunc211(word, this)
}

func SearchFunc211(str string, dict *WordDictionary) bool {
	if len(str) == 0 {
		return dict.end
	}
	if str[0] == '.' {
		for _, v := range dict.children {
			if v != nil && SearchFunc211(str[1:], v) {
				return true
			}
		}
	} else {
		index := str[0] - 'a'
		if dict.children[index] == nil {
			return false
		}
		return SearchFunc211(str[1:], dict.children[index])
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end
