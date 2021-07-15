/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	var reverseNum uint32
	for i := 0; i < 32 && num > 0; i++ {
		reverseNum |= (num & 1) << (31 - i)
		num = num >> 1
	}

	return reverseNum
}

// @lc code=end