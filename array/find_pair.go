package array

import "fmt"

func FindPair(l []int, pair int) [][2]int {
	var match [][2]int
	ll := len(l)
	var counter int = 0
	for i := 0; i < ll; i++ {
		for j := i + 1; j < ll; j++ {
			if l[i]+l[j] == pair {
				match = append(match, [2]int{l[i], l[j]})
			}
			counter++
		}
	}
	fmt.Printf("Normal Counter: %d\n", counter)
	return match
}

func FindPairWithSort(l []int, pair int) [][2]int {
	fmt.Println(l)
	quickSort(l)
	fmt.Println(l)
	ll := len(l)
	var match [][2]int
	var counter int = 0
	for j := ll - 1; j >= 0; j-- {
		for i := 0; i < ll; i++ {
			counter++
			if i == j {
				break
			}
			if l[i]+l[j] == pair {
				match = append(match, [2]int{l[i], l[j]})
			}
		}
	}
	fmt.Printf("SortPair Counter: %d\n", counter)
	return match
}

func FindWithSortV2(l []int, pair int) [][2]int {
	quickSort(l)
	j := len(l) - 1
	var i int = 0
	var match [][2]int
	var counter int = 0
	for i < j {
		counter++
		if l[i]+l[j] == pair {
			match = append(match, [2]int{l[i], l[j]})
			i++
			j--
		} else if l[i]+l[j] > pair {
			j--
		} else {
			i++
		}
	}
	fmt.Printf("SortPairV2 Counter: %d\n", counter)
	return match
}

func quickSort(l []int) {
	partitionSort(l, 0, len(l)-1)
}

func partitionSort(l []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(l, start, end)
	partitionSort(l, start, i-1)
	partitionSort(l, i+1, end)
}

func partition(l []int, start, end int) int {
	pointValue := l[end]
	i := start
	for j := start; j < end; j++ {
		if l[j] > pointValue {
			if i != j {
				l[i], l[j] = l[j], l[i]
			}
			i++
		}
	}
	l[i], l[end] = l[end], l[i]
	return i
}
