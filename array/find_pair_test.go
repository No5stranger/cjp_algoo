package array

import (
	"fmt"
	"testing"
)

func TestFindPair(t *testing.T) {
	var l []int = []int{1, 5, 2, 8, 6, 3, 9}
	match := FindPair(l, 10)
	fmt.Println(match)
	anotherMatch := FindPairWithSort(l, 10)
	fmt.Println(anotherMatch)
	thirdMatch := FindWithSortV2(l, 10)
	fmt.Println(thirdMatch)
}
