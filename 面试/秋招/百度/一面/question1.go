package main

import "fmt"

func question1(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	var result = make([][]int, 0)
	backing(nums, []int{}, &result)
	return result
}

func backing(nums []int, cur []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, cur)
		return
	}
	for i := 0; i < len(nums); i++ {
		cur = append(cur, nums[i])
		temp := make([]int, i)
		copy(temp, nums[:i])
		temp = append(temp, nums[i+1:]...)
		backing(temp, cur, result)
		cur = cur[:len(cur)-1]
	}
}

type test struct {
	name1 int
	name2 string
}

func main() {
	////fmt.Println(question1([]int{1, 2}))
	//fmt.Println(question1([]int{1, 2, 3}))
	////fmt.Println(question1([]int{1, 2, 3, 4}))
	////fmt.Println(question1([]int{1, 2, 3, 4, 5}))
	//fmt.Printf("\n\n\n")
	////fmt.Println(len(question1([]int{1, 2, 3, 4, 5})))
	//
	//fmt.Println(question1([]int{}))
	temp := new(test)
	temp = &test{name1: 1, name2: "qiubing"}
	fmt.Println(temp)
	fmt.Println(*temp)
}
