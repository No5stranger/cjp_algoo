package heap

import "fmt"

type Heap struct {
	// slice to store heap data
	l []int
	// capcity
	n     int
	count int
}

func NewHeap(capcity int) *Heap {
	heap := &Heap{}
	heap.n = capcity
	heap.l = make([]int, capcity+1)
	heap.count = 0
	return heap
}

func (heap *Heap) Insert(data int) {
	if heap.count >= heap.n {
		return
	}
	heap.count++
	heap.l[heap.count] = data
	i := heap.count
	// parent index alway is half of node index
	parent := i / 2
	for parent > 0 && heap.l[parent] < heap.l[i] {
		swap(heap.l, i, parent)
		i = parent
		parent = parent / 2
	}
}

func (heap *Heap) RemoveMax() {
	if heap.count == 0 {
		return
	}
	swap(heap.l, 1, heap.count)
	heap.count--
	heapifyUpToDown(heap.l, heap.count)
}

func (heap *Heap) Print() {
	for i := 0; i < heap.count; i++ {
		fmt.Println(heap.l[i])
	}
}

func heapifyUpToDown(l []int, count int) {
	for i := 1; i <= count/2; {
		maxIndex := i
		if l[i] < l[i*2] {
			maxIndex = i * 2
		}
		if i*2+1 <= count && l[maxIndex] < l[i*2+1] {
			maxIndex = i*2 + 1
		}
		if maxIndex == i {
			break
		}
		swap(l, i, maxIndex)
		i = maxIndex
	}
}

func swap(l []int, i, j int) {
	l[i], l[j] = l[j], l[i]
}
