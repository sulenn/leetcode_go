package checkIfExist

import "fmt"

//https://leetcode-cn.com/contest/weekly-contest-175/problems/check-if-n-and-its-double-exist/

//借助 hashmap，正向遍历和反向遍历两种可能
func checkIfExist(arr []int) bool {
	numMap := make(map[int]bool)
	length := len(arr)
	for i:=0; i<length; i++{  // 正向，0<=i<j<=length-1，arr[i]*2==arr[j]
		if _, ok := numMap[arr[i]*2]; ok {
			return true
		} else {
			numMap[arr[i]] = true
		}
	}
	numMap = make(map[int]bool)
	for i:= length-1; i>=0; i-- {   // 反向,arr[j]*2==arr[i]
		if _, ok := numMap[arr[i]*2]; ok {
			return true
		} else {
			numMap[arr[i]] = true
		}
	}
	return false
}

func main() {
	fmt.Println(checkIfExist([]int {10,2,5,3}))
	fmt.Println(checkIfExist([]int {7,1,14,11}))
	fmt.Println(checkIfExist([]int {3,1,7,11}))
}
