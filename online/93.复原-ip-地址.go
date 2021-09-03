package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}

	ans = []string{}
	backtrack(&s, 0, "", make([]string, 0, 4))

	return ans
}

// var ips map[string]struct{}
var ans []string

func backtrack(s *string, start int, str string, ip []string) {
	if len(ip) > 4 {
		return
	}
	if start > len(*s) {
		return
	}

	if start == len(*s) && len(ip) == 4 {
		ans = append(ans, strings.Join(ip, `.`))
		return
	}

	for i := start; i < len(*s); i++ {
		str += string((*s)[i])
		if len(str) > 1 && strings.HasPrefix(str, `0`) {
			break
		}
		n, _ := strconv.Atoi(str)
		if n > 255 {
			break
		}
		ip = append(ip, str)
		backtrack(s, i+1, "", ip)
		ip = ip[:len(ip)-1]
	}

}

// @lc code=end

func main() {
	fmt.Println(restoreIpAddresses(`25525511135`))
}

// 给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
//
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
//
//
//
// 示例 1：
//
// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]
// 示例 2：
//
// 输入：s = "0000"
// 输出：["0.0.0.0"]
// 示例 3：
//
// 输入：s = "1111"
// 输出：["1.1.1.1"]
// 示例 4：
//
// 输入：s = "010010"
// 输出：["0.10.0.10","0.100.1.0"]
// 示例 5：
//
// 输入：s = "101023"
// 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
