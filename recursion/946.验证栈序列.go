package main

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

// 示例 1：
// 输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
// 输出：true
// 解释：我们可以按以下顺序执行：
// push(1), push(2), push(3), push(4), pop() -> 4,
// push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1

// 示例 2：
// 输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
// 输出：false
// 解释：1 不能在 2 之前弹出。
