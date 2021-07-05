package main

import (
	"fmt"
)

func findErrorNums(nums []int) []int {
	dup := -1
	for i := 0; i < len(nums); i++ {
		index := abs(nums[i]) - 1
		if nums[index] < 0 {
			dup = abs(nums[i])
		} else {
			nums[index] *= -1
		}
	}
	miss := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			miss = i + 1
		}
	}

	return []int{dup, miss}
}

func abs(num int) int {
	if num > 0 {
		return num
	}

	return num * -1
}

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
}
