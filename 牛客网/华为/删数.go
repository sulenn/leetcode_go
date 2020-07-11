package main

import (
	"fmt"
	"io"
)

//约瑟夫环问题
func main() {
	for {
		input := 0
		_, err := fmt.Scan(&input)
		if err == io.EOF {
			break
		}
		result := 0
		for i := 2; i <= input; i++ {
			result = (result + 3) % i
		}
		fmt.Println(result)
	}
}

////用链表存放从 0 - (n-1) 的值
//func main() {
//	for {
//		input := 0
//		_, err := fmt.Scan(&input)
//		if err == io.EOF {
//			break
//		}
//		l := list.List{}
//		for i:=0; i<input; i++ {
//			l.PushBack(i)
//		}
//		for l.Len() != 1 {
//			l.MoveToBack(l.Front())
//			l.MoveToBack(l.Front())
//			l.Remove(l.Front())
//		}
//		fmt.Println(l.Front().Value)
//	}
//}

//用数组存放从 0 - (n-1) 的值
//func main() {
//	for {
//		input := 0
//		_, err := fmt.Scan(&input)
//		if err == io.EOF {
//			break
//		}
//		arr:=[]int {}
//		for i:=0; i<input; i++ {
//			arr = append(arr, i)
//		}
//		index := 0
//		for len(arr) != 1 {
//			for i:=0; i<2; i++ {
//				if index == len(arr) - 1 {
//					index = 0
//				} else {
//					index++
//				}
//			}
//			arr = append(arr[:index], arr[index+1:]...)
//			if index == len(arr) {
//				index = 0
//			}
//		}
//		fmt.Println(arr[0])
//	}
//}
