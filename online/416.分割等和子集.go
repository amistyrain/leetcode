/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2
	dp := make([][]bool, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, sum+1)
	}
	for i := 1; i <= len(nums); i++ {
		dp[i][0] = true
	}
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= sum; j++ {
			if j-nums[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[len(nums)][sum]
}

// @lc code=end

