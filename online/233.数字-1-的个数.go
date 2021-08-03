package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=233 lang=golang
 *
 * [233] 数字 1 的个数
 */

// @lc code=start
func countDigitOne(n int) int {
	p := 0
	for int(math.Pow10(p)) <= n {
		p++
	}
	tmp := make([]int, p)
	for i := 0; i < p; i++ {
		tmp[i] = int(math.Pow10(i))
	}
	// fmt.Println(tmp)
	oneCount := 0
	for i := 1; i <= n; i++ {
		m := i
		for j := len(tmp) - 1; j >= 0; j-- {
			if m < tmp[j] {
				continue
			}
			if m/tmp[j] == 1 {
				oneCount++
			}
			m %= tmp[j]
		}
	}

	return oneCount
}

func main() {
	fmt.Println(countDigitOne(10))
	fmt.Println(countDigitOne(100))
	fmt.Println(countDigitOne(1000))
	fmt.Println(countDigitOne(10000))
	fmt.Println(countDigitOne(20000))
	fmt.Println(countDigitOne(30000))
	fmt.Println(countDigitOne(100000))
	fmt.Println(countDigitOne(1000000))

	fmt.Println(countDigitOne(9))
	fmt.Println(countDigitOne(99))
	fmt.Println(countDigitOne(999))
	fmt.Println(countDigitOne(9999))
	fmt.Println(countDigitOne(99999))

	// fmt.Println(countDigitOne(13))
	// fmt.Println(countDigitOne(824883294))
}

//

// 11
// 1-9 1
// 10-99 10 + 8
// 10-999 100 + 100 + 100

// @lc code=end

// 0000 0001
// 0000 1100
// 0011 0100
