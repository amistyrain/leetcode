package main

import "fmt"

// 暴力
func trap(height []int) int {
	res := 0

	for i := 0; i < len(height); i++ {
		lmax, rmax := 0, 0
		for j := i; j < len(height); j++ {
			rmax = max(rmax, height[j])
		}

		for j := i; j >= 0; j-- {
			lmax = max(lmax, height[j])
		}
		if min(lmax, rmax)-height[i] > 0 {
			res += min(lmax, rmax) - height[i]
		}
	}

	return res
}

func trap2(height []int) int {
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
	fmt.Println(trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap2([]int{4, 2, 0, 3, 2, 5}))
}
