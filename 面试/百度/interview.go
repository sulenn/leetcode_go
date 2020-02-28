package 百度

import (
	"fmt"
	"sync"
)

//反转数组
//func reverse(arr []int) []int {
//	pre := 0
//	last := len(arr) - 1
//	for pre < last {
//		arr[pre], arr[last] = arr[last], arr[pre]
//		pre++
//		last--
//	}
//	return arr
//}

func add(mu *sync.WaitGroup, num []int, ch chan int){
	result := 0
	for _, value := range num {
		result += value
	}
	mu.Done()
	ch <- result
}

//fun 和 add 为并发求和
func fun(num []int) int {
	if len(num) == 0 {return 0}
	mu := &sync.WaitGroup{}
	mu.Add(2)
	sum1 := make(chan int, 1)
	sum2 := make(chan int, 1)
	go add(mu, num[:len(num)/2], sum1)
	go add(mu, num[len(num)/2:], sum2)
	sum3 := <- sum1
	sum4 := <- sum2
	mu.Wait()
	return sum3+sum4
}

func main()  {
	fmt.Println(fun([]int {1,2,3,4,5,6,7,8}))
}

//type Node struct {
//	value int
//	Next *Node
//}

//合并链表
//func fun(l1 *Node, l2 *Node) *Node {
//	if l1==nil {return l2}
//	if l2==nil {return l1}
//	head := &Node{0,nil}
//	cur := head
//	for l1 != nil && l2 != nil {
//		if l1.value > l2.value {
//			cur.Next = l2
//			l2 = l2.Next
//		} else {
//			cur.Next = l1
//			l1 = l1.Next
//		}
//		cur = cur.Next
//	}
//	if l1!=nil {
//		cur.Next = l1
//	} else {
//		cur.Next = l2
//	}
//	return head.Next
//}
//
//func print(l *Node)  {
//	for l != nil {
//		println(l.value)
//		l = l.Next
//	}
//}
