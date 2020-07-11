package main

import "fmt"

//用数组表示的二叉树，可以这样表达： i的子节点下标为 2*i + 1 和 2 * i + 2.   i的父节点下标为 (i-1)/2。
//对于数组长度length，大于length/2的下标一定没有子节点.
func Max_heap(arr []int) []int {
	length := len(arr)
	for i := length/2 - 1; i >= 0; i-- { // 注意完全二叉树种只有从 length/2 - 1 开始的节点才有子树
		build(arr, i, length)
	}
	return arr
}

func build(arr []int, i int, length int) {
	a := 2*i + 1 //  i节点的左右节点为：2*i + 1 和 2 * i + 2
	b := 2*i + 2
	for { // 循环
		if a >= length {
			break
		} // 当前节点不存在左子树
		if b < length && arr[b] > arr[a] {
			a, b = b, a
		}
		if arr[a] > arr[i] {
			arr[i], arr[a] = arr[a], arr[i]
			i = a
		} else {
			break
		}
		a = 2*i + 1
		b = 2*i + 2
	}
}

func print(arr []int) { // 打印最大堆
	length := len(arr) - 1
	for ; length >= 0; length-- {
		fmt.Println(arr[0])
		arr[0], arr[length] = arr[length], arr[0]
		build(arr, 0, length)
	}
}

func main() {
	fmt.Println(Max_heap([]int{1, 5, 3, 6, 7}))
	print(Max_heap([]int{1, 5, 3, 6, 7}))
}
