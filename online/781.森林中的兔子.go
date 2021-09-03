package main

import "fmt"

/*
 * @lc app=leetcode.cn id=781 lang=golang
 *
 * [781] 森林中的兔子
 */

// @lc code=start
func numRabbits(answers []int) int {
	if len(answers) == 0 {
		return 0
	}

	numsMap := make(map[int]int)
	ans := 0
	for i := 0; i < len(answers); i++ {
		if answers[i] == 0 {
			ans += 1
			continue
		}
		if val, ok := numsMap[answers[i]]; !ok {
			numsMap[answers[i]] = answers[i]
			ans += answers[i] + 1
		} else {
			if val > 1 {
				numsMap[answers[i]] = val - 1
			} else {
				delete(numsMap, answers[i])
			}
		}
	}

	return ans
}

// @lc code=end

func main() {
	fmt.Println(numRabbits([]int{0, 0, 1, 1, 1}))
	fmt.Println(numRabbits([]int{10, 10, 10}))
}

// 森林中，每个兔子都有颜色。其中一些兔子（可能是全部）告诉你还有多少其他的兔子和自己有相同的颜色。我们将这些回答放在 answers 数组里。
//
// 返回森林中兔子的最少数量。
//
// 示例:
// 输入: answers = [1, 1, 2]
// 输出: 5
// 解释:
// 两只回答了 "1" 的兔子可能有相同的颜色，设为红色。
// 之后回答了 "2" 的兔子不会是红色，否则他们的回答会相互矛盾。
// 设回答了 "2" 的兔子为蓝色。
// 此外，森林中还应有另外 2 只蓝色兔子的回答没有包含在数组中。
// 因此森林中兔子的最少数量是 5: 3 只回答的和 2 只没有回答的。
//
// 输入: answers = [10, 10, 10]
// 输出: 11
//
// 输入: answers = []
// 输出: 0
