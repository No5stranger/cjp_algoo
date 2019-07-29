package leetcode

import "math"

func AveragePolygonPerimeterByStep(polygon [][2]int, k int) [][2]int {
	var perimeter int = 0
	var pointer, prePointer, nextPointer [2]int
	for i := 1; i <= len(polygon); i++ {
		if i == len(polygon) {
			pointer = polygon[0]
			prePointer = polygon[i-1]
		} else {
			pointer = polygon[i]
			prePointer = polygon[i-1]
		}
		perimeter += distance(pointer, prePointer)
	}
	average := perimeter / k
	var count int = 0
	var averPointers [][2]int
	for count < k {
		dis := average * count
		for j := 0; j < len(polygon); j++ {
			if j == len(polygon)-1 {
				pointer = polygon[j]
				nextPointer = polygon[0]
			} else {
				pointer = polygon[j]
				nextPointer = polygon[j+1]
			}
			pointDistance := distance(pointer, nextPointer)
			if pointDistance >= dis {
				var x, y int
				if pointer[0] == nextPointer[0] {
					if pointer[1] > nextPointer[1] {
						y = pointer[1] - dis
					} else {
						y = pointer[1] + dis
					}
					averPointers = append(averPointers, [2]int{pointer[0], y})
				} else {
					if pointer[0] > nextPointer[0] {
						x = pointer[0] - dis
					} else {
						x = pointer[0] + dis
					}
					averPointers = append(averPointers, [2]int{x, pointer[1]})
				}
				count++
				break
			} else {
				dis -= pointDistance
			}
		}
	}
	return averPointers
}

func distance(a, b [2]int) int {
	return int(math.Abs(float64(a[0]-b[0]))) + int(math.Abs(float64(a[1]-b[1])))
}
