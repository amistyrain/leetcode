package main

import (
	"fmt"
	"math"
	"strconv"
)

func isRectangleCover(rectangles [][]int) bool {
	X1, Y1 := math.MaxInt64, math.MaxInt64
	X2, Y2 := math.MinInt64, math.MinInt64

	points := make(map[string]struct{})
	allAres := 0

	for i := 0; i < len(rectangles); i++ {
		x1 := rectangles[i][0]
		x2 := rectangles[i][2]
		y1 := rectangles[i][1]
		y2 := rectangles[i][3]

		X1, Y1 = min(X1, x1), min(Y1, y1)
		X2, Y2 = max(X2, x2), max(Y2, y2)
		allAres += (y2 - y1) * (x2 - x1)

		fourPoint := []string{
			pointStr(x1, y1),
			pointStr(x1, y2),
			pointStr(x2, y1),
			pointStr(x2, y2),
		}
		for j := 0; j < 4; j++ {
			if _, ok := points[fourPoint[j]]; ok {
				delete(points, fourPoint[j])
			} else {
				points[fourPoint[j]] = struct{}{}
			}
		}
	}
	if len(points) != 4 {
		return false
	}
	if allAres != (Y2-Y1)*(X2-X1) {
		return false
	}
	if _, ok := points[pointStr(X1, Y1)]; !ok {
		return false
	}
	if _, ok := points[pointStr(X1, Y2)]; !ok {
		return false
	}
	if _, ok := points[pointStr(X2, Y1)]; !ok {
		return false
	}
	if _, ok := points[pointStr(X2, Y2)]; !ok {
		return false
	}

	return true
}

func pointStr(x, y int) string {
	return strconv.Itoa(x) + `:` + strconv.Itoa(y)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(isRectangleCover([][]int{
		{1, 1, 3, 3},
		{3, 1, 4, 2},
		{3, 2, 4, 4},
		{1, 3, 2, 4},
		{2, 3, 3, 4},
	}))
}
