package main

import (
	"fmt"
	"io"
	"math"
	"strconv"
)

func main() {
	var nums int  // 多少组函数调用
	for {
		_, err := fmt.Scan(&nums)
		if err == io.EOF {break}
		arr := make([]int, nums)  // 每一组的调用关系
		for i:=0; i<nums; i++ {
			fmt.Scan(&arr[i])
		}
		funcs := make([][]int,nums+1)   // nums + 1
		for i:=1; i<nums+1; i++ {   // 对每组函数进行初始化
			temp := make([]int, arr[i-1]+1)
			fmt.Scan(&temp[0])  // 轮空每一行的下标，该下标对应数组的行
			fmt.Scan(&temp[0])
			for j:=1; j<arr[i-1]+1; j++ {
				fmt.Scan(&temp[j])
			}
			funcs[i] = temp
		}
		fmt.Println(solution(funcs))
	}
}

var max int
var flag1 bool  // 递归调用
var flag2 bool  // 存在未给出调用栈大小的函数

func solution(funcs [][]int) string {
	max = math.MinInt64
	flag1 = false
	flag2 =false
	for i:=1; i<len(funcs); i++ {
		dfs(funcs,i,0,[]int {i})
	}
	if flag1 && flag2 {return "R" +"\r\n" + "NA"}
	if flag1 {return "R"}
	if flag2 {return "NA"}
	return strconv.Itoa(max)
}

func dfs(funcs [][]int, r int, sum int, nums []int){
	sum += funcs[r][0]
	if sum > max {
		max = sum
	}
	for i:=1; i<len(funcs[r]); i++ {
		if judge(funcs[r][i], nums) {
			flag1 = true
			return
		}
		if funcs[r][i] <= 0 || funcs[r][i] >= len(funcs) {
			flag2 = true
			return
		}
		nums = append(nums,funcs[r][i])
		dfs(funcs, funcs[r][i], sum, nums)
		nums = nums[:len(nums)-1]
	}
}

func judge(num int, nums []int) bool {
	for i:=0 ;i<len(nums); i++ {
		if num == nums[i] {return true}
	}
	return false
}
//5 2 3 1 0 1
//1 20 2 3
//2 30 3 4 5
//3 50 4
//4 60
//5 80 9