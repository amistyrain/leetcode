package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getMinLength([][]int{{0, 0}, {0, 2}, {1, 1}, {10, 10}}))
	fmt.Println(getMinLength([][]int{{88, 76}, {12, 57}, {59, 83}, {12, 87}, {77, 0}}))
}

func getMinLength(pearls [][]int) int {
	result := math.MaxInt64
	for i := 0; i < len(pearls)-2; i++ {
		for j := i + 1; j < len(pearls)-1; j++ {
			ij := abs(pearls[j][1]-pearls[i][1]) + abs(pearls[j][0]-pearls[i][0])
			fmt.Println(ij)

			for k := j + 1; k < len(pearls); k++ {
				result = min(result, min(ij+abs(pearls[j][1]-pearls[k][1])+abs(pearls[j][0]-pearls[k][0]), ij+abs(pearls[i][1]-pearls[k][1])+abs(pearls[i][0]-pearls[k][0])))
				fmt.Println(result)
				// result = min(result, )
				// fmt.Println(result)
			}
		}
	}

	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
