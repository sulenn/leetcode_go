package main

//https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/

type CQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() CQueue {
	return CQueue{[]int {},[]int {}}
}

func (this *CQueue) AppendTail(value int)  {
	this.stack1 = append(this.stack1, value)
}


func (this *CQueue) DeleteHead() int {
	if len(this.stack2) == 0 {
		if len(this.stack1) == 0{
			return -1
		} else {
			for len(this.stack1) != 0 {
				this.stack2 = append(this.stack2, this.stack1[len(this.stack1)-1])
				this.stack1 = this.stack1[:len(this.stack1)-1]
			}
		}
	}
	result := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return result
}
