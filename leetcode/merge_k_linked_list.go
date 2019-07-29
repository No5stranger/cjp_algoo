package leetcode

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{val: val}
}

func MergeKLinkedList(l []*ListNode) *ListNode {
	var nodes []int
	var nodeMap map[int][]*ListNode = make(map[int][]*ListNode)
	for i := 0; i < len(l); i++ {
		node := l[i]
		for node != nil {
			keyNode, ok := nodeMap[node.val]
			if !ok {
				nodes = append(nodes, node.val)
				var keyNode []*ListNode = []*ListNode{node}
				nodeMap[node.val] = keyNode
			} else {
				keyNode = append(keyNode, node)
				nodeMap[node.val] = keyNode
			}
			node = node.next
		}
	}
	SelectSort(nodes)
	fmt.Println(nodes)
	var head, pre *ListNode = nil, nil
	for i := 0; i < len(nodes); i++ {
		iNodes := nodeMap[nodes[i]]
		for _, iNode := range iNodes {
			// must reset node next point
			iNode.next = nil
			if head == nil {
				head = iNode
			}
			if pre != nil {
				pre.next = iNode
			}
			pre = iNode
		}
	}
	return head
}

func SelectSort(l []int) {
	for i := 0; i < len(l); i++ {
		minValue := l[i]
		minIndex := i
		for j := i + 1; j < len(l); j++ {
			if l[j] < minValue {
				minValue = l[j]
				minIndex = j
			}
		}
		l[minIndex], l[i] = l[i], l[minIndex]
	}
}
