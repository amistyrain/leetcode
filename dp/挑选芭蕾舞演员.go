package main

import "fmt"

func TeamNums(height []int) int {
	n := len(height)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if height[i] > height[j] && height[j] > height[k] {
					count++
					continue
				}
				if height[i] < height[j] && height[j] < height[k] {
					count++
					continue
				}
			}
		}
	}

	return count
}

func main() {
	fmt.Println(TeamNums([]int{1, 5, 3, 2, 4}))
}
