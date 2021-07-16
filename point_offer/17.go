package main

import (
	"fmt"
	"math"
)

func printNumbers(n int) []int {
	length := int(math.Pow10(n))
	res := make([]int, 0, length)
	for i := 1; i < length; i++ {
		res = append(res, i)
	}

	return res
}

// 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

// 示例 1:

// 输入: n = 1
// 输出: [1,2,3,4,5,6,7,8,9]

func main() {
	fmt.Println(printNumbers(3))
}
