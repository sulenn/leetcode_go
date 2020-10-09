package main

//https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/

type MaxQueue struct {
	q      []int // 数组模拟队列
	maxArr []int // 最大值
}

func Constructor() MaxQueue {
	return MaxQueue{[]int{}, []int{}}
}

func (this *MaxQueue) Max_value() int {
	if len(this.maxArr) != 0 {
		return this.maxArr[0]
	}
	return -1
}

func (this *MaxQueue) Push_back(value int) {
	this.q = append(this.q, value)
	this.maxArr = append(this.maxArr, value)
	for i := len(this.maxArr) - 2; i >= 0; i-- { // 更换最新插入数据之前已有的最大值
		if this.maxArr[i] < this.maxArr[len(this.maxArr)-1] {
			this.maxArr[i] = this.maxArr[len(this.maxArr)-1]
		}
	}
}

func (this *MaxQueue) Pop_front() int {
	if len(this.q) <= 0 {
		return -1
	}
	result := this.q[0]
	this.q = this.q[1:]
	this.maxArr = this.maxArr[1:]
	return result
}
