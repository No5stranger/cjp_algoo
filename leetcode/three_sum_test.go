package leetcode

import "testing"

func TestThreeSum(t *testing.T) {
	var l []int = []int{-1, 0, 1, 2, -1, -4}
	match := ThreeSum(l)
	t.Log(match)
}
