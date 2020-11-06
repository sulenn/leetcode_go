package main

import (
	"math/rand"
)

//https://leetcode-cn.com/problems/random-pick-with-weight/

type Solution struct {
	WeightSum []int
	Sum       int
	Length    int
}

func Constructor(w []int) Solution {
	weightSum := make([]int, len(w))
	sum := 0
	for i := 0; i < len(w); i++ {
		sum += w[i]
		weightSum[i] = sum
	}
	return Solution{
		WeightSum: weightSum,
		Sum:       sum,
		Length:    len(w),
	}
}

func (this *Solution) PickIndex() int {
	random := rand.Intn(this.Sum) + 1 // rand.Intn return [0, n)
	low := 0
	high := this.Length - 1
	for high > low {
		mid := low + (high-low)/2
		if this.WeightSum[mid] == random {
			return mid
		}
		if this.WeightSum[mid] > random {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
