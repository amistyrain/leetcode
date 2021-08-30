package main

import "container/list"

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
type LRUCache struct {
	size   int
	cache  map[int]*list.Element
	helper *list.List
}

type kv struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		size:   capacity,
		cache:  make(map[int]*list.Element, capacity),
		helper: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.cache[key]
	if !ok {
		return -1
	}

	v := val.Value.(*kv)
	this.helper.MoveToFront(val)

	return v.value
}

func (this *LRUCache) Put(key int, value int) {
	val, ok := this.cache[key]
	if ok {
		val.Value.(*kv).value = value
		this.helper.MoveToFront(val)
		return
	}

	if len(this.cache) == this.size {
		val := this.helper.Back()
		delete(this.cache, val.Value.(*kv).key)
		this.helper.Remove(val)
	}

	v := this.helper.PushFront(&kv{key: key, value: value})
	this.cache[key] = v

	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
