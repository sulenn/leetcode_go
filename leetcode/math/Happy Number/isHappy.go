package main

import "fmt"

//https://leetcode-cn.com/problems/happy-number/

func squareAdd(n int) int {
	sum := 0
	for n != 0 {
		sum += (n%10)*(n%10)
		n /= 10
	}
	return sum
}

//参考：https://leetcode-cn.com/problems/happy-number/solution/shi-yong-kuai-man-zhi-zhen-si-xiang-zhao-chu-xun-h/
//快慢指针的方法
func isHappy(n int) bool {
	slow, fast := n, n
	slow = squareAdd(slow)
	fast = squareAdd(fast)
	fast = squareAdd(fast)
	for slow != fast {
		slow = squareAdd(slow)
		fast = squareAdd(fast)
		fast = squareAdd(fast)
	}
	return slow == 1
}

func main() {
	fmt.Println(isHappy(19))
}
