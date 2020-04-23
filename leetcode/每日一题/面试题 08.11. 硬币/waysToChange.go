package main

import "fmt"

//https://leetcode-cn.com/problems/coin-lcci/

func waysToChange(n int) int {
	arr := make([]int,n+1)
	arr[0] = 1
	coins := []int{1,5,10,25}
	for i:=0; i<4; i++ {
		for j:=coins[i]; j<=n; j++ {
			arr[j] = (arr[j] + arr[j-coins[i]]) % 1000000007
		}
	}
	return arr[n]
}

func main() {
	fmt.Println(waysToChange(2))
	fmt.Println(waysToChange(4))
	fmt.Println(waysToChange(5))
	fmt.Println(waysToChange(10))
}