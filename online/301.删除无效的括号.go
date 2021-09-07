package main

/*
 * @lc app=leetcode.cn id=301 lang=golang
 *
 * [301] 删除无效的括号
 */

// @lc code=start
func removeInvalidParentheses(s string) []string {
	ans = []string{}
	path = []byte{}
	set = make(map[string]string)
	str := []byte(s)
	leftRemove, rightRemove := 0, 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			leftRemove++
		} else if str[i] == ')' {
			if leftRemove == 0 {
				rightRemove++
			} else {
				leftRemove--
			}
		}
	}
	dfs(str, 0, 0, 0, leftRemove, rightRemove)
	for i := range set {
		ans = append(ans, i)
	}
	return ans
}

var ans []string
var path []byte
var set map[string]string

func dfs(s []byte, start int, leftcount, rightcount int, leftRemove, rightRemove int) {
	if start == len(s) {
		if leftRemove == 0 && rightRemove == 0 {
			set[string(path)] = ""
			//ans = append(ans, string(path))
		}
		return
	}
	char := s[start]
	if char == '(' && leftRemove > 0 {
		dfs(s, start+1, leftcount, rightcount, leftRemove-1, rightRemove)
	}
	if char == ')' && rightRemove > 0 {
		dfs(s, start+1, leftcount, rightcount, leftRemove, rightRemove-1)
	}

	path = append(path, char)

	if char != '(' && char != ')' {
		dfs(s, start+1, leftcount, rightcount, leftRemove, rightRemove)
	} else if char == '(' {
		dfs(s, start+1, leftcount+1, rightcount, leftRemove, rightRemove)
	} else if rightcount < leftcount {
		dfs(s, start+1, leftcount, rightcount+1, leftRemove, rightRemove)
	}

	path = path[:len(path)-1]
}

// @lc code=end
