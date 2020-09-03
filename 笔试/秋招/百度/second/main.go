package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var T int
	fmt.Scan(&T)
	var n, m int
	var k int
	for w := 0; w < T; w++ {
		fmt.Scan(&n, &m)
		fmt.Scan(&k)
		arr := make([][2]int, k)
		for j := 0; j < k; j++ {
			fmt.Scan(&arr[j][0], &arr[j][1])
		}
		sort.Slice(arr, func(x, y int) bool {
			return arr[x][1] < arr[y][1]
		})

		for i := 1; i < m; i++ {
			fmt.Scan(&k)
			tempArr := make([][2]int, k)
			for j := 0; j < k; j++ {
				fmt.Scan(&tempArr[j][0], &tempArr[j][1])
			}
			sort.Slice(tempArr, func(x, y int) bool {
				return tempArr[x][1] < tempArr[y][1]
			})
			arr = merge(arr, tempArr)
		}
		sum, nums := cal(arr)
		fmt.Println(sum)
		var str string
		for i := 0; i < len(nums); i++ {
			str += strconv.Itoa(nums[i]) + " "
		}
		str = strings.TrimRight(str, " ")
		fmt.Println(str)
	}
}

func merge(arr1 [][2]int, arr2 [][2]int) [][2]int {
	p1 := 0
	p2 := 0
	length1 := len(arr1)
	length2 := len(arr2)
	arr := make([][2]int, 0)
	left := 0
	right := 0
	for p1 < length1 && p2 < length2 {
		if arr1[p1][1] < arr2[p2][0] {
			p1++
			continue
		}
		if arr1[p1][0] > arr2[p2][1] {
			p2++
			continue
		}
		if arr1[p1][0] <= arr2[p2][0] {
			left = arr2[p2][0]
		} else {
			left = arr1[p1][0]
		}
		if arr1[p1][1] <= arr2[p2][1] {
			right = arr1[p1][1]
			p1++
		} else {
			right = arr2[p2][1]
			p2++
		}
		arr = append(arr, [2]int{left, right})
	}
	return arr
}

func cal(arr [][2]int) (int, []int) {
	sum := 0
	nums := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		left := arr[i][0]
		right := arr[i][1]
		for left <= right {
			sum++
			nums = append(nums, left)
			left++
		}
	}
	return sum, nums
}
