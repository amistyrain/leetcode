package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	lIdx := bSearch(nums, target-1)
	rIdx := bSearch(nums, target) - 1
	if lIdx <= rIdx && nums[lIdx] == target {
		return []int{lIdx, rIdx}
	}

	return []int{-1, -1}
}

func bSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := len(nums)
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}

	return ans
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	lIdx := bSearch(nums, target-1)
	rIdx := bSearch(nums, target) - 1
	if lIdx <= rIdx && nums[lIdx] == target {
		return rIdx - lIdx + 1
	}

	return 0

}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 8, 8, 8, 10}, 8))
}
