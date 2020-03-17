package main

type NumArray struct {
	sums []int
}


func Constructor(nums []int) NumArray {
	num := NumArray{make([]int,len(nums))}
	for i, v:= range nums {
		if i == 0 {
			num.sums[i]=v
		} else {
			num.sums[i] = num.sums[i-1] + v
		}
	}
	return num
}


func (this *NumArray) SumRange(i int, j int) int {
	if j < i || j > len(this.sums) || i < 0 {return 0}
	if i == 0 {return this.sums[j]}
	return this.sums[j] - this.sums[i-1]
}
