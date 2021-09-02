package main

/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left < right {
		p := partition(nums, left, right)
		quickSort(nums, left, p-1)
		quickSort(nums, p+1, right)
	}

	return
}

func partition(nums []int, left, right int) int {
	if left >= right {
		return left
	}
	p := nums[left]
	i, j := left, right
	for i < j {
		for nums[j] >= p && i < j {
			j--
		}
		for nums[i] <= p && i < j {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[left], nums[j] = nums[j], nums[left]

	return j
}

// @lc code=end
