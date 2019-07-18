package topk

import "fmt"

func FindBySelectSort(l []int, k int) []int {
	ll := len(l)
	var maxIndex, maxValue int
	for i := 0; i < ll; i++ {
		maxIndex = i
		maxValue = l[i]
		for j := i + 1; j < ll; j++ {
			if maxValue < l[j] {
				maxIndex = j
				maxValue = l[j]
			}
		}
		l[maxIndex], l[i] = l[i], l[maxIndex]
	}
	fmt.Println(l)
	return l[0:k]
}

func FindByQuickSort(l []int, k int) []int {
	QuickSort(l)
	return l[0:k]
}

func QuickSort(l []int) {
	partitionSort(l, 0, len(l)-1)
}

func partitionSort(l []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(l, start, end)
	partitionSort(l, 0, i-1)
	partitionSort(l, i+1, end)
}

func partition(l []int, start, end int) int {
	tmp := l[end]
	i := start
	for j := start; j < end; j++ {
		if l[j] > tmp {
			if i != j {
				l[i], l[j] = l[j], l[i]
			}
			i++
		}
	}
	l[end], l[i] = l[i], l[end]
	return i
}

func FindByHeap(l []int, k int) []int {
	l = append([]int{0}, l...)
	buildHeap(l)
	return l[1 : k+1]
}

func buildHeap(l []int) {
	ll := len(l)
	for i := ll / 2; i >= 1; i-- {
		heaptifyUpToDown(l, i, ll-1)
	}
}

func heaptifyUpToDown(l []int, top, count int) {
	for i := top; i <= count/2; {
		maxIndex := i
		if l[i] < l[i*2] {
			maxIndex = i * 2
		}
		if i*2+1 <= count && l[maxIndex] < l[i*2+1] {
			maxIndex = i*2 + 1
		}
		if i == maxIndex {
			break
		}
		l[maxIndex], l[i] = l[i], l[maxIndex]
		i = maxIndex
	}
}
