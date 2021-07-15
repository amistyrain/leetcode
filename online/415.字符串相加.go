/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	res := ``
	n1 := len(num1) - 1
	n2 := len(num2) - 1
	carry := 0
	for n1 >= 0 || n2 >= 0 || carry != 0 {
		var x, y int
		if n1 >= 0 {
			x = int(num1[n1] - '0')
		}
		if n2 >= 0 {
			y = int(num2[n2] - '0')
		}
		sum := x + y + carry
		carry = sum / 10
		res = strconv.Itoa(sum%10) + res
		n1--
		n2--
	}

	return res
}

// @lc code=end

