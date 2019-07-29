package leetcode

import "fmt"

func StoreWater(l []int) int {
	// 思路： 从底想上逐层计算
	var storage int = 0
	var tmpStorage int = 0
	var left, right, point int
	top := findTheTop(l)
	for f := 1; f <= top; f++ {
		tmpStorage = 0
		for j := 1; j < len(l)-1; j++ {
			point = l[j]
			left = l[j-1]
			right = l[j+1]
			// f:1
			// case: [2, 1, 2] 不需要处理
			if left-f >= 0 && point-f >= 0 && right-f >= 0 {
				continue
			}
			// case: [1, 0, 1] 储水+1
			if point < left && point < right {
				storage++
			}
			// case: [1, 0, 0, 0, 1]  point index: 3
			if point == left && point < right && tmpStorage > 0 {
				tmpStorage++
				storage += tmpStorage
				tmpStorage = 0
			}
			// case: [1, 0, 0, 0, 1]  point index: 1
			if point < left && point == right {
				tmpStorage++
			}
			// case: [1, 0, 0, 0, 1]  point index: 2
			if point == left && point == right && tmpStorage > 0 {
				tmpStorage++
			}
			// 计算后填平, 防止再储水
			if left < f {
				l[j-1] = f
			}
			// 最后一睹墙也需要填平
			if j == len(l)-2 && right < f {
				l[j+1] = f
			}
			fmt.Println(j)
			fmt.Println(storage)
			fmt.Println(l)
		}
		fmt.Printf("floor:%d water:%d\n", f, storage)
	}
	return storage
}

func findTheTop(l []int) int {
	var top int = 0
	for i := 0; i < len(l); i++ {
		if l[i] > top {
			top = l[i]
		}
	}
	return top
}

func OnStorageWater(l []int) int {
	// 思路：寻找左右最高两墙，两墙之间，均低于两墙的可认为可储水
	var storage int = 0
	var left_index int = 0
	var right_index int = len(l) - 1
	var left_max int = l[left_index]
	var right_max int = l[right_index]
	for left_index < right_index {
		if left_max < right_max {
			left_index++
			if left_max < l[left_index] {
				left_max = l[left_index]
			} else {
				storage += left_max - l[left_index]
			}
		} else {
			right_index--
			if right_max < l[right_index] {
				right_max = l[right_index]
			} else {
				storage += right_max - l[right_index]
			}
		}
	}
	return storage
}
