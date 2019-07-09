package sort

import "math/rand"

func MakeRandomSlice(l int) []int {
	var rl []int
	for len(rl) < l {
		rl = append(rl, rand.Intn(100))
	}
	return rl
}

func BubbleSort(l []int, n int) {
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		var flag bool = false
		for j := 0; j < n-1; j++ {
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func InsertionSort(l []int, n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		value := l[i]
		j := i - 1
		for ; j >= 0; j-- {
			if l[j] > value {
				l[j+1] = l[j]
			} else {
				break
			}
		}
		l[j+1] = value
	}
}

func SelectSort(l []int, n int) {
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if l[j] < l[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			l[i], l[minIndex] = l[minIndex], l[i]
		}
	}
}
