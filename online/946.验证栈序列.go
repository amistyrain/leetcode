/*
 * @lc app=leetcode.cn id=946 lang=golang
 *
 * [946] 验证栈序列
 */

// @lc code=start
func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	i, j := 0, 0

	for i < len(pushed) && j < len(pushed) {
		stack = append(stack, pushed[i])
		i++
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			j++
			stack = stack[:len(stack)-1]
		}
	}

	return i == j
}

// @lc code=end

