package sort

func QuickSort(l []int) {
	sperateSort(l, 0, len(l)-1)
}

func sperateSort(l []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(l, start, end)
	sperateSort(l, start, i-1)
	sperateSort(l, i+1, end)
}

func partition(l []int, start, end int) int {
	pointValue := l[end]
	var i = start
	for j := start; j < end; j++ {
		if l[j] < pointValue {
			if i != j {
				l[i], l[j] = l[j], l[i]
			}
			i++
		}
	}
	l[i], l[end] = l[end], l[i]
	return i
}
