package main

import (
	"fmt"
	"sort")

// https://leetcode-cn.com/problems/sort-integers-by-the-number-of-1-bits/

//参考：https://leetcode-cn.com/problems/sort-integers-by-the-number-of-1-bits/solution/javaliang-ci-xun-huan-da-bai-100-by-yourtion/
func sortByBits(arr []int) []int {
	newArr:=make([]int, len(arr))
	count := 0
	num := 0
	for i:=0; i<len(arr); i++ {
		num = arr[i]
		count = 0
		for num != 0 {
			if num & 1 != 0 {
				count++
			}
			num = num >> 1
		}
		newArr[i] = count*1000000+arr[i]
	}
	sort.Ints(newArr)
	for i:=0; i<len(newArr); i++ {
		newArr[i] %= 1000000
	}
	return newArr
}

func main() {
	fmt.Println(sortByBits([]int{0,1,2,3,4,5,6,7,8}))

}