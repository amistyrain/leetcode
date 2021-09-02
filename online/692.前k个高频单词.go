package main

import (
	"container/heap"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 */

// @lc code=start

type word struct {
	w string
	n int
}

type wordsHeap []word

func (ws wordsHeap) Len() int {
	return len(ws)
}

func (ws wordsHeap) Less(i, j int) bool {
	return ws[i].n > ws[j].n || (ws[i].n == ws[j].n && ws[i].w < ws[j].w)
}

func (ws wordsHeap) Swap(i, j int) {
	ws[i], ws[j] = ws[j], ws[i]
}

func (ws *wordsHeap) Pop() interface{} {
	old := *ws
	n := len(old)
	x := old[n-1]
	*ws = old[:n-1]
	return x
}

func (ws *wordsHeap) Push(val interface{}) {
	*ws = append(*ws, val.(word))
}

func topKFrequent(words []string, k int) []string {
	wordsMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		wordsMap[words[i]]++
	}
	wordHeap := make(wordsHeap, len(wordsMap))
	for w, count := range wordsMap {
		wordHeap = append(wordHeap, word{w: w, n: count})
	}
	heap.Init(&wordHeap)
	ans := make([]string, 0, k)
	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(&wordHeap).(word).w)
	}

	return ans
}

// @lc code=end

// 示例 1：
//
// 输入: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
// 输出: ["i", "love"]
// 解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
//     注意，按字母顺序 "i" 在 "love" 之前。
//
//
// 示例 2：
//
// 输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
// 输出: ["the", "is", "sunny", "day"]
// 解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
//     出现次数依次为 4, 3, 2 和 1 次。

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
}
