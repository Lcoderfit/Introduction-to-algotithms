package randAlgorithm

import "math/rand"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution1 struct {
	head *ListNode
}


/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor1(head *ListNode) Solution {
	return Solution1{head: head}
}


/** Returns a random node's value. */
func (this *Solution1) GetRandom() int {
	cnt := 0
	p := this.head
	ret := 0
	for p != nil {
		cnt++
		if rand.Intn(cnt) == 0 {
			ret = p.Val
		}
		p = p.Next
	}
	return ret
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */