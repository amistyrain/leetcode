package main

import "fmt"

func subarraySum(nums []int, k int) int {
	preSums := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		preSums[i+1] = preSums[i] + nums[i]
	}

	count := 0
	for i := 1; i <= len(nums); i++ {
		for j := 0; j < i; j++ {
			if k == preSums[i]-preSums[j] {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
}
