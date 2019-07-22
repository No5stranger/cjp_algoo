package leetcode

import "fmt"

type Stack struct {
	data []int
	top  int
	min  int
}

func NewStack() *Stack {
	return &Stack{data: make([]int, 0), top: -1, min: 0}
}

func (stack *Stack) IsEmpty() bool {
	return stack.top < 0
}

func (stack *Stack) Push(d int) {
	if stack.top < 0 {
		stack.top = 0
	} else {
		stack.top++
	}
	stack.data = append(stack.data, d-stack.min)
	if d < stack.min {
		stack.min = d
	}
}

func (stack *Stack) Pop() int {
	if stack.IsEmpty() {
		return 0
	}
	var vd int
	sd := stack.data[stack.top]
	if sd < 0 {
		vd = stack.min
		stack.min = stack.min - sd
	} else {
		vd = sd + stack.min
	}
	stack.top--
	return vd
}

func (stack *Stack) Top() int {
	if stack.IsEmpty() {
		return 0
	}
	var vd int
	sd := stack.data[stack.top]
	if sd < 0 {
		vd = stack.min
	} else {
		vd = sd + stack.min
	}
	return vd
}

func (stack *Stack) GetMin() int {
	return stack.min
}

func (stack *Stack) Print() {
	if stack.IsEmpty() {
		fmt.Println("empty stack")
	}
	for i := 0; i <= stack.top; i++ {
		fmt.Printf("stack i:%d v:%d\n", i, stack.data[i])
	}
}
