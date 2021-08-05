package main

import (
	"fmt"
	"math"
)

func getFraction(a float32) []int {
	fractNum := float32(math.MaxFloat32)
	result := make([]int, 2)
	for i := 1; i < 200; i++ {
		for j := i + 1; j < 200; j++ {
			if fractNum > abs(float32(i)/float32(j)-a) {
				fractNum = abs(float32(float32(i)/float32(j) - a))
				result[0] = i
				result[1] = j
			}
			if fractNum == abs(float32(i)/float32(j)-a) {
				if i+j < result[0]+result[1] {
					result[0] = i
					result[1] = j
				}
			}
		}
	}

	return result
}

func abs(a float32) float32 {
	if a < 0 {
		return -a
	}
	return a
}

func work(difficulty []int, step []int, ability []int) int {
	sum := 0
	for i := 0; i < len(ability); i++ {
		tmp := 0
		idx := -1
		for j := 0; j < len(difficulty); j++ {
			if difficulty[j] <= ability[i] {
				tmp = max(tmp, difficulty[j])
				idx = j
			}
		}
		if idx >= 0 {
			tmp2 := 0
			for k := 0; k <= idx; k++ {
				if ability[idx] >= difficulty[k] {
					tmp2 = max(tmp2, step[k])
				}
			}
			sum += tmp2

		}
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(getFraction(0.333333333))
	fmt.Println(work([]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7}))
	fmt.Println(work([]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, []int{3, 4, 2, 1, 2}))
}
