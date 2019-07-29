package leetcode

import "fmt"

type Node struct {
	k    string
	v    string
	pre  *Node
	next *Node
}

type LRU struct {
	head  *Node
	tail  *Node
	kvMap map[string]*Node
	maxL  int
}

func NewNode(k, v string) *Node {
	return &Node{k: k, v: v}
}

func NewLRU(maxL int) *LRU {
	var head *Node = NewNode("", "")
	var tail *Node = NewNode("", "")
	head.next = tail
	tail.pre = head
	return &LRU{head: head, tail: tail, kvMap: make(map[string]*Node), maxL: maxL}
}

func (lru *LRU) Put(k, v string) *Node {
	if node, ok := lru.kvMap[k]; ok {
		unlinkNode(node)
	}
	var node *Node = NewNode(k, v)
	lru.AppendHead(node)
	lru.kvMap[k] = node
	if len(lru.kvMap) > lru.maxL {
		tail := lru.RemoveTail()
		delete(lru.kvMap, tail.k)
	}
	return node
}

func (lru *LRU) AppendHead(node *Node) {
	var next *Node = lru.head.next
	lru.head.next = node
	node.pre = lru.head
	node.next = next
	next.pre = node
}

func (lru *LRU) RemoveTail() *Node {
	var tail *Node = lru.tail.pre
	var pre *Node = tail.pre
	lru.tail.pre = pre
	pre.next = lru.tail
	tail.next = nil
	tail.pre = nil
	return tail
}

func (lru *LRU) Print() {
	var next *Node = lru.head.next
	fmt.Println("=== lru start ===")
	for next != nil && next.next != nil {
		fmt.Printf("key:%s value:%s\n", next.k, next.v)
		next = next.next
	}
	fmt.Println("=== lru end ===")
}

func unlinkNode(node *Node) {
	var pre *Node = node.pre
	var next *Node = node.next
	pre.next = next
	next.pre = pre
	node.pre = nil
	node.next = nil
}
