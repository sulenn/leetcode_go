package main

func merge(arr []int, start int, mid int, end int) {
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
		} else {
			arr[end] = right[rightLen-1]
			rightLen--
		}
		end--
	}
	for rightLen > 0 { //  归并的右边数组有剩余
		arr[end] = right[rightLen-1]
		rightLen--
		end--
	}
}

//归并排序
func sort(arr []int, start int, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	sort(arr, start, mid)
	sort(arr, mid+1, end)
	merge(arr, start, mid, end)
}
