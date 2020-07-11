package main

//https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/

func merge(arr []int, start int, mid int, end int) int {
	result := 0
	leftLen := mid + 1 - start
	left := make([]int, leftLen) // 归并的左边数组
	copy(left, arr[start:mid+1]) // 深拷贝
	rightLen := end - mid
	right := make([]int, rightLen) // 归并的右边数组
	copy(right, arr[mid+1:end+1])  // 深拷贝
	for leftLen > 0 && rightLen > 0 {
		if right[rightLen-1] < left[leftLen-1] {
			arr[end] = left[leftLen-1]
			leftLen--
			result += rightLen // 逆序对累加
		} else {
			arr[end] = right[rightLen-1]
			rightLen--
		}
		end--
	}
	for rightLen > 0 { //  归并的右边数组右剩余
		arr[end] = right[rightLen-1]
		rightLen--
		end--
	}
	return result
}

//归并排序
func sort(arr []int, start int, end int) int {
	if start >= end {
		return 0
	}
	mid := (start + end) / 2
	left := sort(arr, start, mid)
	right := sort(arr, mid+1, end)
	result := merge(arr, start, mid, end)
	return left + right + result
}

func reversePairs(nums []int) int {
	return sort(nums, 0, len(nums)-1)
}
