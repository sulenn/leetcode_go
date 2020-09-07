package main

import (
	"container/list"
	"fmt"
)

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */

type LRUCache struct {
	ll        *list.List
	cache     map[int]*list.Element
	maxLength int
	curLength int
}

type Entry struct {
	key   int
	value int
}

func (l *LRUCache) set(key, value int) {
	if ele, ok := l.cache[key]; ok {
		ele.Value.(*Entry).value = value
		l.ll.MoveToBack(ele)
		return
	}
	if l.curLength == l.maxLength {
		l.curLength--
		ele := l.ll.Front()
		l.ll.Remove(ele)
		delete(l.cache, ele.Value.(*Entry).key)
	}
	entry := &Entry{key: key, value: value}
	//ele := &list.Element{Value: entry}
	ele := l.ll.PushBack(entry)
	l.cache[key] = ele
	l.curLength++
}

func (l *LRUCache) get(key int) (int, bool) {
	if ele, ok := l.cache[key]; ok {
		l.ll.MoveToBack(ele)
		return ele.Value.(*Entry).value, true
	}
	return -1, false
}

func NewLRUCache(k int) *LRUCache {
	return &LRUCache{
		ll:        list.New(),
		cache:     make(map[int]*list.Element),
		maxLength: k,
		curLength: 0,
	}
}

func LRU(operators [][]int, k int) []int {
	// write code here
	lrucache := NewLRUCache(k)
	result := make([]int, 0)
	for i := 0; i < len(operators); i++ {
		if operators[i][0] == 1 {
			lrucache.set(operators[i][1], operators[i][2])
		}
		if operators[i][0] == 2 {
			value, _ := lrucache.get(operators[i][1])
			result = append(result, value)
		}
	}
	return result
}

func main() {
	//fmt.Println(LRU([][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}, 3))
	nums := []int{2, 1, 4, 5, 3, 7, 8, 6}
	l := 0
	r := len(nums) - 1
	partition(nums, l, r)
	fmt.Println(nums)
}

func partition(nums []int, left, right int) {
	if left >= right {
		return
	}
	less := left
	v := nums[left]
	for i := left + 1; i <= right; i++ {
		if nums[i] > v {
			less++
			nums[less], nums[i] = nums[i], nums[less]
		}
	}
	nums[less], nums[left] = nums[left], nums[less]
	partition(nums, left, less-1)
	partition(nums, less+1, right)
}
