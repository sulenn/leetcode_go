package main

import (
	"container/list"
	"fmt"
)

// 参考：Go语言实现LRU算法
type LRUCache struct {
	Capacity int
	DList    *list.List
	CacheMap map[interface{}]*list.Element
}

// 之所以需要 CacheNode{Key, Value}，因为当 set 的时候，发现队列满了需要删掉队头元素，此时从队头元素需要获取到 key。
// 而队列中的元素都是 list.Element 对象，且需要和 CacheMap 中 value 值保持一直。所以 CacheNode 中包括 key 和 value
type CacheNode struct {
	Key   interface{}
	Value interface{}
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Capacity: capacity,
		DList:    list.New(),
		CacheMap: make(map[interface{}]*list.Element),
	}
}

func (lru *LRUCache) Size() int {
	return lru.DList.Len()
}

func (lru *LRUCache) Get(key interface{}) interface{} {
	if Element, ok := lru.CacheMap[key]; ok {
		lru.DList.MoveToBack(Element)
		return Element.Value.(*CacheNode).Value
	}
	return nil
}

func (lru *LRUCache) Set(k, v interface{}) {
	if Element, ok := lru.CacheMap[k]; ok { // 更新旧值
		lru.DList.MoveToBack(Element)
		Element.Value.(*CacheNode).Value = v
		return
	}
	if lru.DList.Len() == lru.Capacity { // 已满, 删掉一个元素
		frontElem := lru.DList.Front()
		delete(lru.CacheMap, frontElem.Value.(*CacheNode).Key)
		lru.DList.Remove(frontElem)
	}
	element := lru.DList.PushBack(&CacheNode{k, v})
	lru.CacheMap[k] = element
}

func (lru *LRUCache) Remove(k interface{}) bool {

	if pElement, ok := lru.CacheMap[k]; ok {
		cacheNode := pElement.Value.(*CacheNode)
		delete(lru.CacheMap, cacheNode.Key)
		lru.DList.Remove(pElement)
		return true
	}
	return false
}

func main() {
	lru := NewLRUCache(3)

	lru.Set(10, "value1")
	lru.Set(20, "value2")
	lru.Set(30, "value3")
	lru.Set(10, "value4")
	lru.Set(50, "value5")

	fmt.Println("LRU Size:", lru.Size())
	v := lru.Get(30)
	fmt.Println("Get(30) : ", v)

	if lru.Remove(30) {
		fmt.Println("Remove(30) : true ")
	} else {
		fmt.Println("Remove(30) : false ")
	}
	fmt.Println("LRU Size:", lru.Size())
}
