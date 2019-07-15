package tree

import "testing"

func TestArrayStack(t *testing.T) {
	as := NewArrayStack()
	for i := 0; i < 8; i++ {
		as.Push(i)
	}
	pop_data := as.Pop()
	t.Log(pop_data)
	as.Print()
}
