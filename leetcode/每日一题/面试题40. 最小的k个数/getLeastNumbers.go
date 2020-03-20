package main

//https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/submissions/

//最大堆
//该题也可以用快排，找到快排点和 k 值相等即可
type myheap struct {
	v []int
}

func (h *myheap) Len() int {
	return len(h.v)
}

func (h *myheap) Less(i,j int) bool {
	return h.v[i] > h.v[j]
}

func (h *myheap) Swap(i,j int) {
	h.v[i], h.v[j] = h.v[j], h.v[i]
}

func (h *myheap) Push(v interface{})  {
	h.v = append(h.v, v.(int))
}

func (h *myheap) Pop() interface{} {
	v := h.v[len(h.v)-1]
	h.v = h.v[:len(h.v)-1]
	return v
}

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int {}
	}
	h := &myheap{}
	heap2.Init(h)
	for _, v:= range arr {
		if h.Len() != k {
			heap2.Push(h, v)
		} else if h.v[0] > v {
			heap2.Pop(h)
			heap2.Push(h, v)
		}
	}
	return h.v
}
