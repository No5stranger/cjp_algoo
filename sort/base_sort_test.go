package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var l []int = MakeRandomSlice(5)
	fmt.Println(l)
	BubbleSort(l, 5)
	fmt.Println(l)
}

func TestInsertionSort(t *testing.T) {
	var l []int = MakeRandomSlice(5)
	fmt.Println(l)
	InsertionSort(l, 5)
	fmt.Println(l)
}

func TestSelectSort(t *testing.T) {
	var l []int = MakeRandomSlice(5)
	fmt.Println(l)
	SelectSort(l, 5)
	fmt.Println(l)
}
