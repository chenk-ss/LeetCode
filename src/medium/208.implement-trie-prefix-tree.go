package medium

/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start
type Trie struct {
	children [26]*Trie
	end      bool
}

func Constructor208() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, v := range word {
		index := v - 'a'
		if cur.children[index] == nil {
			cur.children[index] = &Trie{}
		}
		cur = cur.children[index]
	}
	cur.end = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, v := range word {
		index := v - 'a'
		if cur.children[index] == nil {
			return false
		}
		cur = cur.children[index]
	}
	return cur.end
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, v := range prefix {
		index := v - 'a'
		if cur.children[index] == nil {
			return false
		}
		cur = cur.children[index]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
