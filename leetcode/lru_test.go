package leetcode

import "testing"

func TestLRU(t *testing.T) {
	var seeds map[string]string = map[string]string{"a": "aa", "b": "bb", "c": "cc", "d": "dd"}
	lru := NewLRU(2)
	for k, v := range seeds {
		lru.Put(k, v)
		lru.Print()
	}
}
