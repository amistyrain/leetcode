package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(twoSum([]int{2, 7, 9, 12}, 9))
	fmt.Println(twoSum([]int{-1, 0, 1, 2, -1, -4}, -4))
	// fmt.Println(twoSum([]int{2, 7, 9, 12, 4, 5}, 9))
	// fmt.Println(twoSum([]int{2, 7, 9, 12, 4, 5, 4, 5}, 9))

	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func twoSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	left := 0
	right := len(nums) - 1
	for left < right {
		l := nums[left]
		r := nums[right]
		sum := r + l
		if target < sum {
			for left < right && r == nums[right] {
				right--
			}
		} else if target > sum {
			for left < right && l == nums[left] {
				left++
			}
		} else {
			result = append(result, []int{r, l})
			for left < right && l == nums[left] {
				left++
			}
			for left < right && r == nums[right] {
				right--
			}
		}
	}

	return result
}

func tSum(nums []int, i int, target int) [][]int {
	// sort.Ints(nums)
	result := make([][]int, 0)
	left := i
	right := len(nums) - 1
	for left < right {
		l := nums[left]
		r := nums[right]
		sum := r + l
		if target < sum {
			for left < right && r == nums[right] {
				right--
			}
		} else if target > sum {
			for left < right && l == nums[left] {
				left++
			}
		} else {
			result = append(result, []int{r, l})
			for left < right && l == nums[left] {
				left++
			}
			for left < right && r == nums[right] {
				right--
			}
		}
	}

	return result
}

// func threeSum(nums []int) [][]int {
// 	sort.Ints(nums)
// 	result := make([][]int, 0)
// 	target := 0
// 	for i := 0; i < len(nums); i++ {
// 		tsum := tSum(nums, i+1, target-nums[i])

// 		for _, v := range tsum {
// 			result = append(result, append([]int{nums[i]}, v...))
// 		}
// 		for i < len(nums)-1 && nums[i] == nums[i+1] {
// 			i++
// 		}
// 	}

// 	return result
// }

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := nSum(nums, 3, 0, 0)
	return result
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := nSum(nums, 4, 0, target)
	return result
}

func nSum(nums []int, n int, start int, target int) [][]int {
	left := start
	right := len(nums) - 1
	result := make([][]int, 0)
	if n == 2 {
		for left < right {
			l := nums[left]
			r := nums[right]
			sum := r + l
			if target < sum {
				for left < right && r == nums[right] {
					right--
				}
			} else if target > sum {
				for left < right && l == nums[left] {
					left++
				}
			} else {
				result = append(result, []int{r, l})
				for left < right && l == nums[left] {
					left++
				}
				for left < right && r == nums[right] {
					right--
				}
			}
		}
	} else {
		for i := start; i < len(nums); i++ {
			tsum := nSum(nums, n-1, i+1, target-nums[i])

			for _, v := range tsum {
				result = append(result, append([]int{nums[i]}, v...))
			}
			for i < len(nums)-1 && nums[i] == nums[i+1] {
				i++
			}
		}

	}

	return result
}
