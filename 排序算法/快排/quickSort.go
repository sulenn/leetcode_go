package main

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	point1 := 0            // point1 指向的下标元素值固定不变
	point2 := len(arr) - 1 // 从后往前遍历或者从前往后遍历的下标
	for point1 != point2 {
		if point1 < point2 { // point2 在数组中的位置大于 point1
			if arr[point1] > arr[point2] { // point2 指向的元素小于 point1 指向的元素值
				arr[point1], arr[point2] = arr[point2], arr[point1]
				point1, point2 = point2, point1
				point2++
			} else {
				point2--
			}
		} else { // point2 在数组中的位置小于 point1
			if arr[point1] < arr[point2] { // point2 指向的元素大于 point1 指向的元素值
				arr[point1], arr[point2] = arr[point2], arr[point1]
				point1, point2 = point2, point1
				point2--
			} else {
				point2++
			}
		}
	}
	quickSort(arr[:point1])
	quickSort(arr[point1+1:])
}
