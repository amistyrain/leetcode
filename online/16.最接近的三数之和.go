/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := 0
	minNum := math.MaxInt64
	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			l := nums[left]
			r := nums[right]
			sum := r + l + nums[i]
			if sum == target {
				return target
			}

			if abs(target-sum) < minNum {
				minNum = abs(target - sum)
				res = sum
			}
			if sum > target {
				right--
			} else {
				left++
			}
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}

// @lc code=end

