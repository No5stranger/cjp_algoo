package leetcode

import "fmt"

func FindMaxSubString(s string) int {
	var sPos map[rune]int = make(map[rune]int)
	var start int = 0
	var maxLength int = 0
	for i, ch := range []rune(s) {
		if pos, ok := sPos[ch]; ok && pos >= start {
			start = pos + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		fmt.Printf("i:%d start:%d maxL:%d\n", i, start, maxLength)
		sPos[ch] = i
	}
	return maxLength
}
