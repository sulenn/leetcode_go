package main

import "fmt"

//思路：字典存储key/value值，Caches中存放key操作次数
//同时Caches中key从左到右的顺序满足最不经常使用和最少使用原则

//O(1)的解决方法可以参考：https://leetcode-cn.com/problems/lfu-cache/solution/ha-xi-biao-shuang-xiang-lian-biao-java-by-liweiwei/
//哈希表 + 双向链表

type LFUCache struct {
	Dic      map[int]int
	Caches   []Cache
	Capacity int
}

type Cache struct {
	key   int
	count int
}

func Constructor(capacity int) LFUCache {
	dic := make(map[int]int, 0)
	caches := make([]Cache, 0)
	return LFUCache{Dic: dic, Caches: caches, Capacity: capacity}
}

func (this *LFUCache) Get(key int) int {
	value := 0
	if _, ok := this.Dic[key]; ok {
		value = this.Dic[key]
	} else {
		return -1
	}
	ModifySort(&this.Caches, key)
	return value
}

func (this *LFUCache) Put(key int, value int) {
	if this.Capacity == 0 {
		return
	}
	if _, ok := this.Dic[key]; ok {
		this.Dic[key] = value
		ModifySort(&this.Caches, key)
		return
	}
	if len(this.Caches) == this.Capacity {
		deleteKey := this.Caches[0]
		this.Caches = this.Caches[1:]
		delete(this.Dic, deleteKey.key)
	}
	this.Dic[key] = value
	temp := []Cache{Cache{key, 0}}
	this.Caches = append(temp, this.Caches...)
	ModifySort(&this.Caches, key)
}

func ModifySort(caches *[]Cache, key int) {
	for i := 0; i < len(*caches); i++ {
		if (*caches)[i].key == key { // 找到满足要求的下标
			(*caches)[i].count++
			for j := i + 1; j < len(*caches); j++ {
				if (*caches)[j].count > (*caches)[i].count {
					temp := make([]Cache, len(*caches)-j)
					copy(temp, (*caches)[j:])
					*caches = append((*caches)[:j], (*caches)[i])
					*caches = append(*caches, temp...)
					*caches = append((*caches)[:i], (*caches)[i+1:]...)
					return
				}
				if j == len(*caches)-1 {
					*caches = append(*caches, (*caches)[i])
					*caches = append((*caches)[:i], (*caches)[i+1:]...)
					return
				}
			}
		}
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}
