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

func (this *LinkedList) PrintLinkedList() {
	cur := this.head.next
	fmt.Println(fmt.Sprintf("linkedlist length %d", this.GetLength()))
	for nil != cur {
		fmt.Println(fmt.Sprintf("ndoe:%d", cur.value))
		cur = cur.next
	}
}

func main() {
	testLinkedList := NewLinkedList()
	testLinkedList.InsertToHead(0)
	testLinkedList.PrintLinkedList()
	testLinkedList.InsertToTail(1)
	testLinkedList.PrintLinkedList()
	testLinkedList.InsertToTail(2)
	testLinkedList.PrintLinkedList()
	testLinkedList.Revert()
	testLinkedList.PrintLinkedList()
	testLinkedList.DeleteNodeByValue(1)
	testLinkedList.PrintLinkedList()
}
