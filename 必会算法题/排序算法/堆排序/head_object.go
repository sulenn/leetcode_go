package main

import "fmt"

// 最大堆
type Heap struct {
	Size int
	Nums []int
}

func NewHeap() *Heap {
	return &Heap{Size: -1, Nums: make([]int, 0)} // 注意 size 初始化 -1
}

func (h *Heap) Push(num int) {
	h.Size++
	if h.Size == len(h.Nums) { // 溢出，扩容
		h.Nums = append(h.Nums, 0)
	}
	i := h.Size
	for {
		if i <= 0 {
			break
		}
		parent := (i - 1) / 2
		if h.Nums[parent] >= num {
			break

		}
		h.Nums[i] = h.Nums[parent]
		i = parent
	}
	h.Nums[i] = num
}

func (h *Heap) Pop() int {
	if h.Size < 0 {
		return -1
	}
	//h.Size--
	result := h.Nums[0]
	h.Nums[0] = h.Nums[h.Size]
	h.Size--
	i := 0
	for {
		left := i*2 + 1
		right := i*2 + 2
		if right <= h.Size && h.Nums[left] < h.Nums[right] {
			h.Nums[left], h.Nums[right] = h.Nums[right], h.Nums[left]
		}
		if left <= h.Size && h.Nums[i] < h.Nums[left] {
			h.Nums[i], h.Nums[left] = h.Nums[left], h.Nums[i]
		} else {
			break
		}
		i = left
	}
	return result
}

func (h *Heap) Print() {
	fmt.Printf("h.Nums: %v, h.Size: %v\n", h.Nums, h.Size)
}

func main() {
	h := NewHeap()
	h.Print()

	h.Push(3)
	h.Push(6)
	h.Push(7)
	h.Push(27)
	h.Push(1)
	h.Push(2)
	h.Push(3)
	h.Print()

	fmt.Println(h.Pop())
	h.Print()
	fmt.Println(h.Pop())
	h.Print()
	fmt.Println(h.Pop())
	h.Print()
	fmt.Println(h.Pop())
	h.Print()
	fmt.Println(h.Pop())
	h.Print()
	fmt.Println(h.Pop())
	h.Print()
	fmt.Println(h.Pop())
	h.Print()
	fmt.Println(h.Pop())
	h.Print()

	h.Push(20)
	h.Print()
	h.Push(15)
	h.Print()
	h.Push(30)
	h.Print()
}
