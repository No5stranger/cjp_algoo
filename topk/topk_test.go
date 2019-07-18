package topk

import (
	"fmt"
	"testing"
)

func TestFindBySelectSort(t *testing.T) {
	var l []int = []int{5, 9, 12, 3, 7}
	match := FindBySelectSort(l, 2)
	fmt.Println(match)
}

func TestFindByQuickSort(t *testing.T) {
	var l []int = []int{5, 9, 12, 3, 7}
	match := FindByQuickSort(l, 2)
	fmt.Println(match)
}

func TestFindByHeap(t *testing.T) {
	var l []int = []int{5, 9, 12, 3, 7}
	match := FindByHeap(l, 2)
	fmt.Println(match)
}
