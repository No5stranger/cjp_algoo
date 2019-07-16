package heap

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestHeap(t *testing.T) {
	heap := NewHeap(9)
	var record map[int]int
	record = make(map[int]int)
	var i int = 0
	var data int
	for i < 9 {
		data = rand.Intn(100)
		_, ok := record[data]
		if !ok {
			record[data] = data
			heap.Insert(data)
			i++
		}
	}
	heap.Print()
	heap.RemoveMax()
	heap.Print()
}

func TestHeapWithSlice(t *testing.T) {
	var l []int = []int{0, 5, 9, 4, 10, 66, 8}
	BuildHeap(l, 6)
	fmt.Println(l)
	Sort(l, 6)
	fmt.Println(l)
}
