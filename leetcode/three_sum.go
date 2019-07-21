package leetcode

import "fmt"

func ThreeSum(l []int) [][]int {
	var match [][]int
	insertSort(l)
	fmt.Println(l)
	for i := 0; i < len(l); i++ {
		if i > 0 && l[i] == l[i-1] {
			continue
		}
		a := l[i]
		j := i + 1
		z := len(l) - 1
		for j < z {
			b := l[j]
			c := l[z]
			if a+b+c > 0 {
				z--
			} else if a+b+c < 0 {
				j++
			} else {
				match = append(match, []int{a, b, c})
				for j < z && l[j] == l[j+1] {
					j++
				}
				for j < z && l[z] == l[z-1] {
					z--
				}
				j++
				z--
			}
		}
	}
	return match
}

func insertSort(l []int) {
	for i := 1; i < len(l); i++ {
		value := l[i]
		j := i - 1
		for ; j >= 0; j-- {
			if l[j] > value {
				l[j], l[j+1] = l[j+1], l[j]
			} else {
				break
			}
		}
		l[j+1] = value
	}
}
