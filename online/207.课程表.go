package main

import "fmt"

/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		valid   = true
		dfs     func(u int)
		result  = []int{}
	)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}

	return valid
}

// @lc code=end

func main() {
	fmt.Println(canFinish(6, [][]int{{5, 3}, {5, 4}, {3, 0}, {3, 1}, {4, 1}, {4, 2}}))
}
