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
	//l := 0
	//r := len(nums) - 1
	//partition(nums, l, r)
	//fmt.Println(nums)
	HeapSort(nums)
	fmt.Println(nums)
	//length := len(nums)
	//for i := 0; i < length; i++ {
	//	fmt.Println(pop(nums))
	//}
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

func heapSort(nums []int) {
	length := len(nums)
	mid := length / 2
	for mid >= 0 {
		curIndex := mid
		for curIndex < length {
			left := curIndex*2 + 1
			if left >= length {
				break
			}
			right := curIndex*2 + 2
			if right < length && nums[left] < nums[right] {
				nums[left], nums[right] = nums[right], nums[left]
			}
			if nums[curIndex] < nums[left] {
				nums[curIndex], nums[left] = nums[left], nums[curIndex]
			} else {
				break
			}
			curIndex = left
		}
		mid--
	}
}

func pop(nums []int) (int, bool) {
	if len(nums) == 0 {
		return -1, false
	}
	result := nums[0]
	length := len(nums)
	nums[0] = nums[length-1]
	nums = nums[:length-1]
	index := 0
	length--
	for index < length {
		left := index*2 + 1
		if left >= length {
			break
		}
		right := index*2 + 2
		if right < length && nums[left] < nums[right] {
			nums[left], nums[right] = nums[right], nums[left]
		}
		if nums[index] < nums[left] {
			nums[index], nums[left] = nums[left], nums[index]
		} else {
			break
		}
		index = left
	}
	return result, true
}

func HeapSort(values []int) {
	buildHeap(values)
	for i := len(values); i > 1; i-- {
		values[0], values[i-1] = values[i-1], values[0]
		adjustHeap(values[:i-1], 0)
		//fmt.Println("the heap is ", values)
	}
}

func buildHeap(values []int) {
	for i := len(values) / 2; i >= 0; i-- { //////一定得从后往前调整，
		adjustHeap(values, i)
	}
}

func adjustHeap(values []int, pos int) { ///////// 调整pos位置的结点
	node := pos
	length := len(values)
	for node < length {
		var child = node*2 + 1
		if child+1 < length && values[child+1] > values[child] {
			child++
		}

		if child < length && values[child] > values[node] {
			values[node], values[child] = values[child], values[node]
			node = child
		} else {
			break
		}
	}
}
