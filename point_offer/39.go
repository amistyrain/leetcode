package main

import "fmt"

func majorityElement(nums []int) int {
	num := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			num = nums[i]
			count = 1
			continue
		}
		if nums[i] == num {
			count++
		} else {
			count--
		}
	}

	return num
}

func main() {
	fmt.Println(majorityElement([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
