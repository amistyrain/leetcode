package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 1 || len(nums) == 0 {
		return len(nums)
	}
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}

	return slow + 1
}

func main() {
	fmt.Println([]int{1, 2, 2, 3, 4, 5, 5, 6})
	fmt.Println(removeDuplicates([]int{1, 2, 2, 3, 4, 5, 5, 6}))
}
