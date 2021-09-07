package main

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &Trie{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}

	return node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for i := 0; i < len(prefix); i++ {
		idx := prefix[i] - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
