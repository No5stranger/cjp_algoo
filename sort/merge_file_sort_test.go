package sort

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
)

func TestSortByMergeFile(t *testing.T) {
	fd, err := os.OpenFile("test_sort_file", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		t.Failed()
	}
	for i := 0; i < 10; i++ {
		fd.WriteString(strconv.Itoa(rand.Intn(1000)))
		fd.WriteString("\n")
	}
	SortByMergeFile()
}
