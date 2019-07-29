package leetcode

import "testing"

func TestAveragePolygonPerimeter(t *testing.T) {
	var polygon [][2]int = [][2]int{{1, 1}, {1, 2}, {2, 2}, {2, 3}, {4, 3}, {4, 1}}
	pointers := AveragePolygonPerimeterByStep(polygon, 5)
	t.Log(pointers)
}
