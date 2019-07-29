package leetcode

import "testing"

func TestFindMaxSubString(t *testing.T) {
	var ss []string = []string{"abcabcbb", "bbbb", "pwwkew", "abcabct"}
	for _, s := range ss {
		maxL := FindMaxSubString(s)
		t.Logf("s:%s max sub string length:%d", s, maxL)
	}
}
