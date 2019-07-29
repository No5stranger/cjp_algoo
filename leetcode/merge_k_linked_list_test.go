package leetcode

import "testing"

func TestMergeKLinkedList(t *testing.T) {
	a := NewListNode(1)
	b := NewListNode(4)
	e := NewListNode(5)
	b.next = e
	a.next = b
	aa := NewListNode(1)
	c := NewListNode(3)
	d := NewListNode(4)
	c.next = d
	aa.next = c
	cc := NewListNode(2)
	f := NewListNode(6)
	cc.next = f
	var l []*ListNode = []*ListNode{a, c, cc}
	head := MergeKLinkedList(l)
	t.Log(head.val)
	for head != nil {
		t.Log(head.val)
		head = head.next
	}
}
