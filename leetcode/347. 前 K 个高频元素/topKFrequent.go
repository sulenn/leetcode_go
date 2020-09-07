package main

import (
	"sort"
)

//https://leetcode-cn.com/problems/top-k-frequent-elements/

type num struct {
	n     int
	count int
}

func topKFrequent(nums []int, k int) []int {
	dic := make(map[int]int)
	for _, v := range nums {
		dic[v]++
	}
	numArr := []num{}
	for k, v := range dic {
		numArr = append(numArr, num{k, v})
	}
	sort.Slice(numArr, func(i, j int) bool {
		return numArr[i].count > numArr[j].count
	})
	result := make([]int, len(numArr))
	for k, v := range numArr {
		result[k] = v.n
	}
	return result[:k]
}
