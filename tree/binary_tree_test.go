package tree

import "testing"

func makeBinaryTree() *BinaryTree {
	bt := NewBinaryTree("a")
	bNode := NewNode("b")
	bNode.left = NewNode("d")
	bNode.right = NewNode("e")
	cNode := NewNode("c")
	cNode.left = NewNode("f")
	cNode.right = NewNode("g")
	bt.root.left = bNode
	bt.root.right = cNode
	return bt
}

func TestBinaryTree_InOrder(t *testing.T) {
	bt := makeBinaryTree()
	bt.InOrderTraverse()
}

func TestBinaryTree_PreOrder(t *testing.T) {
	bt := makeBinaryTree()
	bt.PreOrderTraverse()
}

func TestBinaryTree_PostOrderWithTwoStack(t *testing.T) {
	bt := makeBinaryTree()
	bt.PostOrderTraverseWithTwoStack()
}

func TestRecursiveInOrderTraverse(t *testing.T) {
	bt := makeBinaryTree()
	RecursiveInOrderTraverse(bt.root)
}
