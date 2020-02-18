package main

import "fmt"

//https://leetcode-cn.com/problems/count-primes/
//注意是小于非负整数 n，不包括n

//n 为 499979，超时
//func countPrimes(n int) int {
//	if n < 3 {return 0}
//	primeArray := []int {2}
//	for i:=3; i<n; i++ {
//		flag := true
//		for _,value := range primeArray {
//			if i % value == 0 {
//				flag = false
//				break
//			}
//		}
//		if flag {
//			primeArray = append(primeArray, i)
//		}
//	}
//	return len(primeArray)
//}

//解法参考：https://leetcode-cn.com/problems/count-primes/solution/ru-he-gao-xiao-pan-ding-shai-xuan-su-shu-by-labula/。
//时间和空间复杂度均为 o(n)
func countPrimes(n int) int {
	if n < 3 {return 0}  // 不包括n的非负整数
	primeArray := make([]bool, n)
	for index,_ :=range primeArray[2:] {   // 排除0和1
		primeArray[index+2] = true
	}
	for i:=2; i*i < n; i++ {
		if !primeArray[i] {
			continue
		}
		for j:=2*i; j<n; j+=i {
			primeArray[j]=false
		}
	}
	result := 0
	for _,value := range primeArray {
		if value {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(countPrimes(10))
	//fmt.Println(countPrimes(2))
	//fmt.Println(countPrimes(3))
	//fmt.Println(countPrimes(4))
	//fmt.Println(countPrimes(5))
}
