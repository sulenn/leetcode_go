package main

import (
	"fmt"
	"io"
)

func main() {
	var num int
	var nums [6][2]int
	for {
		_, err := fmt.Scan(&num)
		if err == io.EOF {break}
		var a int
		var b int
		for i:=0; i<num; i++ {
			for j:=0; j<6; j++ {
				fmt.Scan(&a,&b)
				if a < b {
					nums[j][0], nums[j][1] = a, b
				} else {
					nums[j][0], nums[j][1] = b, a
				}
			}
			fmt.Println(solution(nums))
		}
	}
}
//1
//1345 2584
//2584 683
//2584 1345
//683 1345
//683 1345
//2584 683

//1
//1345 1345
//1345 1345
//1345 1345
//1345 1345
//1345 1345
//1345 1345

func solution(nums [6][2]int) string {
	//dic := make(map[int]int)
	var arr = make([][]int,0)
	//for i:=0; i<6; i++ {
	//	if v, ok := dic[nums[i][0]]; ok {
	//		if v != nums[i][1] {
	//			flag := false
	//			for u:=0; u< len(arr);u++ {
	//				if arr[u][0] == nums[i][0] && arr[u][1] == nums[i][1] {
	//					flag = true
	//				}
	//			}
	//			if !flag {
	//				arr = append(arr, []int {nums[i][0], nums[i][1]})
	//			}
	//		}
	//	} else {
	//		dic[nums[i][0]] = nums[i][1]
	//	}
	//}
	arr = append(arr, []int {nums[0][0], nums[0][1]})
	for i:=1; i<6; i++ {
		flag := false
		for j:=0; j<len(arr); j++ {
			if arr[j][0] == nums[i][0] && arr[j][1] == nums[i][1] {
				flag = true
				break
			}
		}
		if !flag {
			arr = append(arr, []int {nums[i][0], nums[i][1]})
		}
	}
	if len(arr) > 3 {return "IMPOSSIBLE"}
	//for k,v := range dic {
	//	arr = append(arr, []int{k,v})
	//}
	for len(arr) < 3 {
		if len(arr) == 2 {
			if arr[0][0] != arr[0][1] {
				arr = append(arr,arr[0])
			} else {
				arr = append(arr,arr[1])
			}
		} else {
			arr = append(arr,arr[0])
			arr = append(arr,arr[0])
		}
	}
	if arr[0][0] == arr[1][0] {   // 第一个数相同
		arr[0][0], arr[1][0] = 0, 0
	}
	if arr[0][0] == arr[2][0] {
		arr[0][0], arr[2][0] = 0, 0
	}
	if arr[1][0] == arr[2][0] {
		arr[1][0], arr[2][0] = 0, 0
	}
	if arr[0][1] == arr[1][1] {  // 第二个数相同
		arr[0][1], arr[1][1] = 0, 0
	}
	if arr[0][1] == arr[2][1] {
		arr[0][1], arr[2][1] = 0, 0
	}
	if arr[1][1] == arr[2][1] {
		arr[1][1], arr[2][1] = 0, 0
	}
	if arr[0][0] == arr[1][1] {  // 第一个数和第二个数相同
		arr[0][0], arr[1][1] = 0, 0
	}
	if arr[0][0] == arr[2][1] {
		arr[0][0], arr[2][1] = 0, 0
	}
	if arr[1][0] == arr[2][1] {
		arr[1][0], arr[2][1] = 0, 0
	}
	if arr[1][0] == arr[0][1] {
		arr[1][0], arr[0][1] = 0, 0
	}
	if arr[2][0] == arr[0][1] {
		arr[2][0], arr[0][1] = 0, 0
	}
	if arr[2][0] == arr[1][1] {
		arr[2][0], arr[1][1] = 0, 0
	}
	for i:=0; i<3; i++ {
		if arr[i][0] != 0 {return "IMPOSSIBLE"}
		if arr[i][1] != 0 {return "IMPOSSIBLE"}
	}
	return "POSSIBLE"
}
