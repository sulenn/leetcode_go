package main

import "fmt"

//https://leetcode-cn.com/problems/ugly-number-ii/

func nthUglyNumber(n int) int {

}

//使用 Hash map 的方法，超时
//func nthUglyNumber(n int) int {
//	uglyMap := make(map[int]bool)
//	result := 0
//	if n <= 0 { return 0 }
//	if n == 1 { return 1 }
//	uglyMap[1]=true
//	n--
//	count := 2
//	for n > 0 {
//		if count%2==0 {
//			if _, ok := uglyMap[count / 2]; ok{
//				uglyMap[count] = true
//				result = count
//				n--
//			}
//		} else if count%3==0 {
//			if _, ok := uglyMap[count / 3]; ok {
//				uglyMap[count] = true
//				result = count
//				n--
//			}
//		} else if count%5==0 {
//			if _, ok := uglyMap[count / 5]; ok {
//				uglyMap[count] = true
//				result = count
//				n--
//			}
//		}
//		count++
//	}
//	return result
//}

func main() {
	//fmt.Println(nthUglyNumber(10))
	fmt.Println(nthUglyNumber(11))
}
