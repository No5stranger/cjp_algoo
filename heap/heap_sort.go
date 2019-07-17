package heap

func Sort(l []int, n int) {
	// move the max element to the tail
	BuildHeap(l, n)
	k := n
	for k > 1 {
		l[1], l[k] = l[k], l[1]
		k--
		heapifyUpToDownWithTop(l, 1, k)
	}
}

func BuildHeap(l []int, n int) {
	// only parent node ned to compare
	for i := n / 2; i >= 1; i-- {
		heapifyUpToDownWithTop(l, i, n)
	}
}

func heapifyUpToDownWithTop(l []int, top, count int) {
	for i := top; i <= count/2; {
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
		l[maxIndex], l[i] = l[i], l[maxIndex]
		i = maxIndex
	}
}
