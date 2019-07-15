package tree

import "fmt"

type ArrayStack struct {
	data []interface{}
	top  int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 5),
		top:  -1,
	}
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.top < 0
}

func (stack *ArrayStack) Push(data interface{}) {
	if stack.top < 0 {
		stack.top = 0
	} else {
		stack.top += 1
	}
	if stack.top > len(stack.data)-1 {
		stack.data = append(stack.data, data)
	} else {
		stack.data[stack.top] = data
	}
}

func (stack *ArrayStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}

	data := stack.data[stack.top]
	stack.top -= 1
	return data
}

func (stack *ArrayStack) Flush() {
	stack.top = -1
}

func (stack *ArrayStack) Print() {
	if stack.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		for i := 0; i < len(stack.data); i++ {
			fmt.Println(stack.data[i])
		}
	}
}
