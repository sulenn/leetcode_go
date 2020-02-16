package main

import "math"

//https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/

type MinStack struct {
	stack []int  // 数字栈用于存储数字
	minStack []int    // 最小值辅助栈
	length int  // 记录栈的长度
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, []int{},0}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	if this.length == 0 {
		this.minStack = append(this.minStack, x)
	} else if this.minStack[this.length-1] < x {
		this.minStack = append(this.minStack, this.minStack[this.length-1])
	} else {
		this.minStack = append(this.minStack, x)
	}
	this.length++
}


func (this *MinStack) Pop()  {
	if this.length != 0 {
		this.stack = this.stack[:this.length-1]
		this.minStack = this.minStack[:this.length-1]
		this.length--
	}
}


func (this *MinStack) Top() int {
	if this.length != 0 {
		return this.stack[this.length-1]
	}
	return math.MaxInt64
}


func (this *MinStack) Min() int {
	if this.length != 0 {
		return this.minStack[this.length-1]
	}
	return math.MaxInt64
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
