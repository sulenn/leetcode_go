package main

import "sort"

//https://leetcode-cn.com/problems/top-k-frequent-elements/

type num struct {
	n     int
	count int
}

func topKFrequent(nums []int, k int) []int {
	dic := make(map[int]int)
	for _, v := range nums { // 用字典记录每个数字出现的次数
		dic[v]++
	}
	numArr := []num{}
	for k, v := range dic { // 将数字和频率放入num结构体切片
		numArr = append(numArr, num{k, v})
	}
	sort.Slice(numArr, func(i, j int) bool { // 从大到小排序切片
		return numArr[i].count > numArr[j].count
	})
	result := make([]int, len(numArr))
	for k, v := range numArr { // 将结构体切片转换成数字切片
		result[k] = v.n
	}
	return result[:k]
}
