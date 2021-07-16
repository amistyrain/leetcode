package main

import (
	"fmt"
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	if len(nums) == 0 {
		return ``
	}

	sort.Slice(nums, func(i, j int) bool {
		if strconv.Itoa(nums[i])+strconv.Itoa(nums[j]) > strconv.Itoa(nums[j])+strconv.Itoa(nums[i]) {
			return false
		}
		return true
	})
	result := ``
	for i := 0; i < len(nums); i++ {
		result += strconv.Itoa(nums[i])
	}

	return result
}

func main() {
	testStr := []string{`330`, `303`}
	sort.Strings(testStr)
	fmt.Println(testStr)

	fmt.Println(minNumber([]int{3, 30, 34, 5, 9}))
}

// 输入: [3,30,34,5,9]
// 输出: "3033459"

// 3 30
// 330
// 303
