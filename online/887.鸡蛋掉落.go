/*
 * @lc app=leetcode.cn id=887 lang=golang
 *
 * [887] 鸡蛋掉落
 */

// @lc code=start
func superEggDrop(k int, n int) int {
	var dp func(int, int) int
	table := make(map[string]int)

	dp = func(k, n int) int {
		if k == 1 {
			return n
		}
		if n == 0 {
			return 0
		}
		if v, ok := table[strconv.Itoa(k)+`:`+strconv.Itoa(n)]; ok {
			return v
		}

		result := math.MaxInt64
		l, r := 1, n
		for l <= r {
			mid := l + (r-l)/2
			broken := dp(k-1, mid-1)
			notBroken := dp(k, n-mid)
			if broken > notBroken {
				r = mid - 1
				result = min(result, broken+1)
			} else {
				l = mid + 1
				result = min(result, notBroken+1)
			}

		}
		table[strconv.Itoa(k)+`:`+strconv.Itoa(n)] = result

		return result
	}

	return dp(k, n)
}

// func superEggDrop(k int, n int) int {
// 	var dp func(int, int) int
// 	table := make(map[string]int)

// 	dp = func(k, n int) int {
// 		if k == 1 {
// 			return n
// 		}
// 		if n == 0 {
// 			return 0
// 		}
// 		if v, ok := table[strconv.Itoa(k)+`:`+strconv.Itoa(n)]; ok {
// 			return v
// 		}

// 		result := math.MaxInt64
// 		for i := 1; i <= n; i++ {
// 			result = min(result,
// 				max(dp(k, n-i), dp(k-1, i-1))+1)
// 			// 碎   dp(k-1,i-1)
// 			// 没碎 dp(k,n-i)
// 			// +1   鸡蛋用了一次
// 		}
// 		table[strconv.Itoa(k)+`:`+strconv.Itoa(n)] = result

// 		return result
// 	}

// 	return dp(k, n)
// }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

