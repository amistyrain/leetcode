package main

import "fmt"

func main() {
	nums := []int{1, 4, 2, 5, 76, 8, 34, 3, 56, 67}
	QuckSort(nums)
	fmt.Println(nums)
}

func QuckSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
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
	if left > right {
		return left
	}
	p := nums[left]
	i, j := left+1, right
	for {
		for nums[i] <= p {
			i++
			fmt.Println(i)
		}
		for nums[j] > p {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		fmt.Println(nums)
	}
	nums[left], nums[i] = nums[i], nums[left]

	fmt.Println(nums)

	return i
}
