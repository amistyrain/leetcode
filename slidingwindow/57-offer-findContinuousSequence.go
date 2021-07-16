package main

import (
	"fmt"
)

func findContinuousSequence(target int) [][]int {
	result := [][]int{}
	left, right := 0, 1
	windowLen := target/2 + 1
	window := make([]int, 0, windowLen)
	for i := 1; i <= windowLen; i++ {
		window = append(window, i)
	}
	for right < len(window) {
		s := sum(window[left : right+1])
		if s < target {
			right++
		} else if s > target {
			left++
		} else {
			result = append(result, window[left:right+1])
			left++
		}
	}

	return result
}

func sum(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += nums[i]
	}

	return ans
}

func main() {
	fmt.Println(findContinuousSequence(9))
}
