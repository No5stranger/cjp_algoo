package leetcode

import "fmt"

type StackV2 struct {
	data []int
	top  int
}

type MinStack struct {
	ds  *StackV2
	ms  *StackV2
	min int
}

func NewStackV2() *StackV2 {
	return &StackV2{data: make([]int, 0), top: -1}
}

func (stack *StackV2) Push(d int) {
	if stack.top < 0 {
		stack.top = 0
	} else {
		stack.top++
	}
	stack.data = append(stack.data, d)
}

func (stack *StackV2) Pop() int {
	if stack.top < 0 {
		return 0
	}
	d := stack.data[stack.top]
	stack.top--
	return d
}

func (stack *StackV2) Top() int {
	if stack.top < 0 {
		return 0
	}
	return stack.data[stack.top]
}

func NewMinStack() *MinStack {
	minStack := NewStackV2()
	return &MinStack{
		ds: NewStackV2(),
		ms: minStack,
	}
}

func (mis *MinStack) Push(d int) {
	if mis.ms.top < 0 {
		mis.ms.Push(d)
		mis.min = d
	}
	mis.ds.Push(d)
	if d <= mis.min {
		mis.ms.Push(d)
		mis.min = d
	}
}

func (mis *MinStack) Pop() int {
	d := mis.ds.Pop()
	if d == mis.min {
		mis.ms.Pop()
		mis.min = mis.ms.Pop()
	}
	fmt.Println(mis.ms)
	fmt.Println(mis.min)
	return d
}

func (mis *MinStack) Top() int {
	return mis.ds.Top()
}

func (mis *MinStack) GetMin() int {
	return mis.min
}
