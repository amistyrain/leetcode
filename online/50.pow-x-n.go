/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := float64(1)

	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		n >>= 1
		x *= x
	}

	return res
}

// @lc code=end

