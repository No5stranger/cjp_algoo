package linkedlist

import "fmt"

type LinkedNode struct {
	next  *LinkedNode
	value interface{}
}

type LinkedList struct {
	head   *LinkedNode
	length uint
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewLinkedNode(0), 0}
}

func (this *LinkedList) GetLength() uint {
	return this.length
}

func NewLinkedNode(v interface{}) *LinkedNode {
	return &LinkedNode{nil, v}
}

func (this *LinkedNode) GetNext() *LinkedNode {
	return this.next
}

func (this *LinkedNode) GetValue() interface{} {
	return this.value
}

func (this *LinkedList) InsertAfter(p *LinkedNode, v interface{}) bool {
	if p == nil {
		return false
	}
	newNode := NewLinkedNode(v)
	pre := p.next
	p.next = newNode
	newNode.next = pre
	this.length++
	return true
}

func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

func (this *LinkedList) DeleteNode(p *LinkedNode) bool {
	if p == nil {
		return false
	}

	cur := this.head.next
	pre := this.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	pre.next = p.next
	p = nil
	this.length--
	return true
}

func (this *LinkedList) DeleteNodeByValue(v interface{}) bool {
	return this.DeleteNode(this.FindByValue(v))
}

func (this *LinkedList) FindByValue(v interface{}) *LinkedNode {
	cur := this.head
	for nil != cur {
		if cur.GetValue() == v {
			break
		}
		cur = cur.next
	}
	return cur
}

func (this *LinkedList) Revert() {
	if this.head.next == nil {
		return
	}
	node := this.RevertByNode(this.head.next)
	this.head.next = node

}

func (this *LinkedList) RevertByNode(head *LinkedNode) *LinkedNode {
	var pre *LinkedNode = nil
	var next *LinkedNode = nil
	for head != nil {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	return pre
}

func (this *LinkedList) RevertSlow() {
	var pre *LinkedNode = nil
	var cur *LinkedNode = this.head.next
	var next *LinkedNode = this.head.next.next
	for i := uint(0); i < this.GetLength(); i++ {
		cur.next = pre
		pre = cur
		cur = next
		if next != nil {
			// not rewrite next if next is the last node
			if next.next != nil {
				next = next.next
			}
		} else {
			next = nil
		}
	}
	this.head.next = cur
}

func (this *LinkedList) FindByIndex(index uint) *LinkedNode {
	if index >= this.GetLength() {
		return nil
	}
	cur := this.head.next
	for i := uint(0); i < index; i++ {
		cur = cur.next
	}
	return cur
}

func (this *LinkedList) GetMidNode() (uint, *LinkedNode) {
	var i uint = 0
	var fast *LinkedNode = this.head.next
	var slow *LinkedNode = this.head.next
	for fast != nil {
		if fast.next != nil {
			fast = fast.next.next
			i++
		} else {
			break
		}
		slow = slow.next
	}
	return i, slow
}

func (this *LinkedList) IsPalindrome() bool {
	_, midNode := this.GetMidNode()
	var pre *LinkedNode = nil
	var next *LinkedNode = nil
	var head = this.head.next
	for head != midNode {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	this.head.next = pre
	var isPalindrome bool = true
	var left *LinkedNode = pre
	var right *LinkedNode
	if this.length%2 != 0 {
		right = midNode.GetNext()
	} else {
		right = midNode
	}
	for nil != left && nil != right {
		if left.GetValue() != right.GetValue() {
			isPalindrome = false
			break
		}
		left = left.GetNext()
		right = right.GetNext()
	}
	head = pre
	pre = midNode
	for head != nil {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	this.head.next = pre
	return isPalindrome
}

func (this *LinkedList) PrintLinkedList() {
	cur := this.head.next
	fmt.Println(fmt.Sprintf("linkedlist length %d", this.GetLength()))
	for nil != cur {
		PrintNode(cur)
		cur = cur.next
	}
}

func PrintNode(node *LinkedNode) {
	fmt.Println(fmt.Sprintf("node:%d", node.value))
}
