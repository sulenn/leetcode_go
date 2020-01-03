package main

//https://leetcode-cn.com/submissions/detail/42096201/

import (
	"fmt"
)

type listNode struct {   // 节点
	key int
	value int
	prev *listNode
	next *listNode
}

type LRUCache struct {   // 主缓存结构
	inMap map[int]*listNode
	capacity int
	head *listNode  // 头指针
	tail *listNode  // 尾指针
}

// 双向链表
func Constructor(capacity int) LRUCache {
	head := &listNode{0, 0,nil,nil}
	tail := &listNode{0,0,nil,nil}
	head.next = tail
	tail.prev = head
	return LRUCache{make(map[int]*listNode, capacity), capacity, head, tail}
}

func (this *LRUCache) Move_node_to_tail(key int) {   // 新访问的元素移动至队尾
	node := this.inMap[key]
	node.prev.next = node.next // 改变节点原本前后节点的指向
	node.next.prev = node.prev

	node.prev = this.tail.prev // 改变节点现在的前后指向
	node.next = this.tail

	this.tail.prev.next = node // 改变尾节点前节点的后指向，以及尾节点自己的前指向
	this.tail.prev = node
}

func (this *LRUCache) Get(key int) int {   // 获取元素
	if node , ok := this.inMap[key]; ok {
		this.Move_node_to_tail(key)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.inMap[key]; ok {  // 节点已存在
		this.Move_node_to_tail(key)
		this.inMap[key].value = value
	} else {
		if len(this.inMap) == this.capacity {   // 缓存已满， 删除队头的元素
			delete(this.inMap, this.head.next.key)
			this.head.next.next.prev = this.head   //改变队头节点的后指向，以及队头元素下个节点的前指向
			this.head.next = this.head.next.next
		}
		node := &listNode{key, value, nil,nil}
		node.next = this.tail  //新节点指针插入队尾前一个结点
		node.prev = this.tail.prev
		this.tail.prev.next = node  //队尾结点指向调整
		this.tail.prev = node
		this.inMap[key] = node   // map 值调整
	}
}

//自己写的垃圾方法，用伪栈的方法
//type LRUCache struct {
//	inMap map[int]int
//	cache []int  // 设定一个伪栈，栈底为 right, 栈顶为 left，中间的元素也可以访问
//	right int
//	left int
//}
//
//
//func Constructor(capacity int) LRUCache {   // 初始化
//	return LRUCache{map[int]int{}, make([]int, capacity), capacity, capacity}
//}
//
//
//func (this *LRUCache) Get(key int) int {
//	if value, ok:= this.inMap[key]; ok {
//		for i:=this.left; i<this.right + 1; i++ {  // 将访问过的元素防止栈顶
//			if this.cache[i] == key {
//				temp := this.cache[i]
//				for j:=i; j > this.left; j-- {  // 移动元素
//					this.cache[j] = this.cache[j-1]
//				}
//				this.cache[this.left] = temp
//				break
//			}
//		}
//		return value
//	} else {
//		return -1
//	}
//}
//
//
//func (this *LRUCache) Put(key int, value int)  {
//	if this.right == len(this.cache) {   // 初始化时，this.right 指向栈底的下一个索引
//		this.right--
//		this.left--
//		this.cache[this.right] = key
//		this.inMap[key] = value
//		return
//	}
//	for i:=this.left; i<this.right+1; i++ {   // 查询新插入的元素是否已存在
//		if this.cache[i] == key {   // 存在则移动该元素至栈顶
//			temp := this.cache[i]
//			for j:=i; j > this.left; j-- {
//				this.cache[j] = this.cache[j-1]
//			}
//			this.cache[this.left] = temp   // 更新map中 value
//			this.inMap[key] = value
//			return
//		}
//	}
//	if this.left > 0 {   // 栈中还存在未被利用的空间
//		this.left--
//		this.cache[this.left] = key
//		this.inMap[key] = value
//		return
//	}
//	delete(this.inMap, this.cache[this.right])   // 栈已满，删掉最没用的元素（栈底）
//	for i:=this.right; i>this.left; i-- {   // 移动元素，更新元素
//		this.cache[i] = this.cache[i-1]
//	}
//	this.cache[this.left] = key
//	this.inMap[key] = value
//}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	//cache := Constructor(2)
	//cache.Put(1, 1);
	//cache.Put(2, 2);
	//fmt.Println(cache.Get(1))
	//cache.Put(3, 3);
	//fmt.Println(cache.Get(2))
	//cache.Put(4, 4);
	//fmt.Println(cache.Get(1))
	//fmt.Println(cache.Get(3))
	//fmt.Println(cache.Get(4))
	cache := Constructor(10)
	cache.Put(10, 13);
	cache.Put(3, 17);
	cache.Put(6, 11);
	cache.Put(10, 5);
	cache.Put(9, 10);
	fmt.Println(cache.Get(13))
	cache.Put(2, 19);
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))
	cache.Put(5, 25);
	fmt.Println(cache.Get(8))
	cache.Put(9, 22);
	cache.Put(5, 5);
	cache.Put(1, 30);
	fmt.Println(cache.Get(11))
	cache.Put(9, 12);
	fmt.Println(cache.Get(7))
	fmt.Println(cache.Get(5))
	fmt.Println(cache.Get(8))
	fmt.Println(cache.Get(9))
	cache.Put(4, 30);
	cache.Put(9, 3);
	fmt.Println(cache.Get(9))
	fmt.Println(cache.Get(10))
	fmt.Println(cache.Get(10))
}
