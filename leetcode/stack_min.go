package leetcode

import "fmt"

type Stack struct {
	data []int64
	top  int64
	min  int64
}

func NewStack() *Stack {
	return &Stack{data: make([]int64, 0), top: -1}
}

func (stack *Stack) IsEmpty() bool {
	return stack.top < 0
}

func (stack *Stack) Push(d int64) {
	if stack.top < 0 {
		stack.top = 0
		stack.min = d
		stack.data = append(stack.data, d)
	} else {
		stack.top++
		stack.data = append(stack.data, d-stack.min)
	}
	if d < stack.min {
		stack.min = d
	}
}

func (stack *Stack) Pop() int64 {
	if stack.IsEmpty() {
		return 0
	}
	var vd int64
	sd := stack.data[stack.top]
	if stack.top == 0 {
		vd = sd
	} else {
		if sd < 0 {
			vd = stack.min
			stack.min = stack.min - sd
		} else {
			vd = sd + stack.min
		}
	}
	stack.top--
	return vd
}

func (stack *Stack) Top() int64 {
	if stack.IsEmpty() {
		return 0
	}
	var vd int64
	sd := stack.data[stack.top]
	if sd < 0 {
		vd = stack.min
	} else {
		vd = sd + stack.min
	}
	return vd
}

func (stack *Stack) GetMin() int64 {
	return stack.min
}

func (stack *Stack) Print() {
	if stack.IsEmpty() {
		fmt.Println("empty stack")
	}
	for i := int64(0); i <= stack.top; i++ {
		fmt.Printf("stack i:%d v:%d\n", i, stack.data[i])
	}
}
