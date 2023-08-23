package medium

/*
 * @lc app=leetcode id=284 lang=golang
 *
 * [284] Peeking Iterator
 */

// @lc code=start
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	arr []int
}

func Constructor284(iter *Iterator) *PeekingIterator {
	res := PeekingIterator{}
	for iter.hasNext() {
		res.arr = append(res.arr, iter.next())
	}
	return &res
}

func (this *PeekingIterator) hasNext() bool {
	return len(this.arr) != 0
}

func (this *PeekingIterator) next() int {
	res := this.arr[0]
	this.arr = this.arr[1:]
	return res
}

func (this *PeekingIterator) peek() int {
	return this.arr[0]
}

// @lc code=end
