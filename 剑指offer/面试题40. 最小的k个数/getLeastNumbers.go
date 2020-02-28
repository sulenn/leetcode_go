package main

import "container/heap"

// 利用 go 中自带的堆结构来做。大顶堆
func getLeastNumbers(arr []int, k int) []int {
	if k > len(arr) || k <= 0 {return []int {}}
	h := &IntHeap{}
	heap.Init(h)
	for i:=0;i<len(arr);i++ {
		if i < k {
			heap.Push(h,arr[i])
		} else if arr[i] < (*h)[0] {
			heap.Pop(h)
			heap.Push(h, arr[i])
		}
	}
	return *h
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i,j int) bool {  // 小于号是小顶堆，大于号是大顶堆
	return h[i] > h[j]
}

func (h IntHeap) Swap(i,j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{})  {   // 注意这里必须是指针
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {   // 注意这里必须是指针
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}