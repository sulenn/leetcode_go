package main

import "fmt"

//

func main() {
	temp1 := 1.5
	temp2 := 1
	fmt.Println(temp1 > temp2)
}

//import (
//	"fmt"
//	"io"
//)
//
//func main() {
//	var n, m int
//	for {
//		_, err := fmt.Scan(&n, &m)
//		if err == io.EOF {
//			break
//		}
//		nums := make([][]int, n)
//		for i:=0; i<n; i++ {
//			nums[i] = make([]int, m)
//			for j:=0; j<m; j++ {
//				fmt.Scan(&nums[i][j])
//			}
//		}
//		first, second :=
//	}
//}
//
//func cal(nums [][]int, first, second int) (int, int) {
//	r := len(nums)
//	c := len(nums[0])
//	flag := make([][]bool, r)
//	for i:=0; i<r; i++ {
//		flag[i] = make([]bool, c)
//	}
//	for i:=0; i<r; i++ {
//		for j:=0; j<c; j++ {
//			count :=
//		}
//	}
//}
//
//func coun(flag [][]bool, nums [][]int, x, y int) int {
//	if flag[x][y] {
//		return 0
//	}
//	result := 1
//
//}
