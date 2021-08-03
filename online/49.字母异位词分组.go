package main

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	strMap := make(map[[26]int][]string)

	for i := 0; i < len(strs); i++ {
		count := [26]int{}
		for j := 0; j < len(strs[i]); j++ {
			count[byte(strs[i][j])-'a']++
		}
		strMap[count] = append(strMap[count], strs[i])
	}
	result := make([][]string, 0, len(strMap))
	for _, v := range strMap {
		result = append(result, v)
	}

	return result
}

// @lc code=end

func main() {

}
