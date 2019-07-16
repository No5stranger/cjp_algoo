package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var l []int = MakeRandomSlice(9)
	fmt.Println(l)
	ml := MergeSort(l)
	fmt.Println(ml)
}
