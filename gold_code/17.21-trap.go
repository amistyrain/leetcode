package main

import "fmt"

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	res := 0
	lmax, rmax := make([]int, len(height)), make([]int, len(height))
	lmax[0] = height[0]
	for i := 1; i < len(height); i++ {
		lmax[i] = max(lmax[i-1], height[i])
	}
	rmax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rmax[i] = max(rmax[i+1], height[i])
	}

	for i := 0; i < len(height); i++ {
		if min(lmax[i], rmax[i])-height[i] > 0 {
			res += min(lmax[i], rmax[i]) - height[i]
		}

	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
