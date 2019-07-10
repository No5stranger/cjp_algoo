package tree

import "fmt"

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(v interface{}) *BinaryTree {
	return &BinaryTree{root: NewNode(v)}
}

func (bt *BinaryTree) InOrderTraverse() {
	var t *Node
	p := bt.root
	s := NewArrayStack()
	for p != nil || !s.IsEmpty() {
		if p != nil {
			s.Push(p)
			p = p.left
		} else {
			t = s.Pop().(*Node)
			fmt.Println(t.data)
			p = t.right
		}
	}
}

func (bt *BinaryTree) PreOrderTraverse() {
	var t *Node
	p := bt.root
	s := NewArrayStack()
	for p != nil || !s.IsEmpty() {
		if p != nil {
			fmt.Println(p.data)
			s.Push(p)
			p = p.left
		} else {
			t = s.Pop().(*Node)
			p = t.right
		}
	}
}

func (bt *BinaryTree) PreOrderTraverse2() {
	var t *Node
	p := bt.root
	s := NewArrayStack()
	for p != nil || !s.IsEmpty() {
		fmt.Println(p.data)
		if p.right != nil {
			s.Push(p)
		}
		if p.left != nil {
			p = p.left
		} else if !s.IsEmpty() {
			t = s.Pop().(*Node)
			p = t.right
		} else {
			p = nil
		}
	}
}

func (bk *BinaryTree) PostOrderTraverse() {
}

func (bk *BinaryTree) PostOrderTraverseWithTwoStack() {
	s1 := NewArrayStack()
	s2 := NewArrayStack()
	s1.Push(bk.root)
	var t *Node
	for !s1.IsEmpty() {
		t = s1.Pop().(*Node)
		s2.Push(t)
		if t.left != nil {
			s1.Push(t.left)
		}
		if t.right != nil {
			s1.Push(t.right)
		}
	}
	for !s2.IsEmpty() {
		t := s2.Pop().(*Node)
		fmt.Println(t.data)
	}
}
