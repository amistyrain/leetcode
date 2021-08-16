package main

/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */

// @lc code=start
func countSubstrings(s string) int {
	n := len(s)
	result := 0
	// 	为什么是2n-1个中心点？
	// 如果回文串是奇数，我们把回文串中心的那个字符叫做中心点，如果回文串是偶数我们就把中间的那两个字符叫做中心点。对于一个长度为n的字符串，我们可以用它的任意一个字符当做中心点，所以中心点的个数是n。我们还可以用它任意挨着的两个字符当做中心点，所以中心点是n-1，总的中心点就是2*n-1
	// for i := 0; i < 2*n-1; i++ {
	// 	l, r := i/2, i/2+i%2
	// 	for l >= 0 && r < n && s[l] == s[r] {
	// 		l--
	// 		r++
	// 		result++
	// 	}
	// }

	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			result++
		}
	}

	for i := 0; i < n-1; i++ {
		l, r := i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			result++
		}
	}

	return result
}

// @lc code=end
