package main

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	p1, p2 := head, head
	for i := 0; i < k-1; i++ {
		if p1.Next != nil {
			p1 = p1.Next
		} else {
			return nil
		}
	}
	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
}

func main() {
	//var temp = make(chan int, 1)
	//count := 0
	//close(temp)
	//for {
	//	select {
	//	case <-temp:
	//		return
	//	default:
	//		time.Sleep(1 * time.Second)
	//		fmt.Println("qiubing1")
	//		if count > 5 {
	//			close(temp)
	//		}
	//		count++
	//	}
	//}
	var i interface{} = 10
	_, ok := i.(int)
	if !ok {
		log.Fatalf("sdjgkajdgka")
	}
}

//type temp interface {
//	print()
//}
//
//type qiubing struct {
//}

//func (q *qiubing) print() {
//	fmt.Println("qiubing")
//}
//
//func inter(in interface{}) {
//	var i interface{} = 10
//	_, ok := i.(string)
//	if !ok {
//		log.Fatalf("sdjgkajdgka")
//	}
//
//}
