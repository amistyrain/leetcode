package main

import "fmt"

func exchange(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	left, right := 0, len(nums)-1
	for left < right {
		for nums[right]%2 == 0 && right > 0 {
			right--
		}
		for nums[left]%2 != 0 && left < len(nums)-1 {
			left++
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	return nums
}

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
	fmt.Println(exchange([]int{1, 2, 3, 4, 6, 7, 8, 10, 33, 11}))
	fmt.Println(exchange([]int{1, 3, 5}))
	fmt.Println(exchange([]int{2, 4, 6}))
}
