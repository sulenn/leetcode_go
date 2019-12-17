package main

import "fmt"

//78. Subsets https://leetcode-cn.com/problems/subsets/solution/hui-su-suan-fa-by-powcai-5/

//迭代
func subsets(nums []int) [][]int {
	//numLen := len(nums)
	var result [][]int
	result = append(result, []int {})
	for _ ,v := range nums {
		for _, v1 := range result{
			result = append(result, append([]int {v}, v1...))
		}
	}
	return result
}

//递归(回溯算法)
//var result [][]int
//var numLen int
//var realNums []int
//
//func subsets(nums []int) [][]int {
//	numLen = len(nums)
//	realNums = nums
//	helper(0, []int {})
//	return result
//}
//
//func helper(i int, tmp []int, )  {
//	result = append(result, tmp)
//	for j := i; j < numLen; j++ {
//		helper(j + 1, append(tmp, realNums[j]))
//	}
//}

func main() {
	//print(subsets([]int {1,2,3}))
	fmt.Println(subsets([]int {1,2,3}))
}


