package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var l []int = MakeRandomSlice(15)
	fmt.Println(l)
	QuickSort(l)
	fmt.Println(l)
}
