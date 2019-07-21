package leetcode

import "testing"

func TestStorageWater(t *testing.T) {
	var l []int = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	water := StoreWater(l)
	t.Logf("storage water:%d", water)
	var l2 []int = []int{2, 5, 1, 2, 3, 4, 7, 7, 6}
	water2 := StoreWater(l2)
	t.Logf("storage water:%d", water2)
}

func TestOnStorageWater(t *testing.T) {
	var l []int = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	water := OnStorageWater(l)
	t.Logf("storage water:%d", water)
	var l2 []int = []int{2, 5, 1, 2, 3, 4, 7, 7, 6}
	water2 := OnStorageWater(l2)
	t.Logf("storage water:%d", water2)
}
