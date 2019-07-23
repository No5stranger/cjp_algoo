package leetcode

import "testing"

func TestMinStack(t *testing.T) {
	stack := NewStack()
	//stack := NewMinStack()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	//stack.Print()
	t.Log(stack.GetMin())
	stack.Pop()
	t.Log(stack.GetMin())
	stack.Pop()
	t.Log(stack.GetMin())
	stack.Pop()
	//stack.Print()
	t.Log(stack.GetMin())
}
