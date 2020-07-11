package main

import "fmt"

//https://leetcode-cn.com/contest/weekly-contest-171/problems/minimum-flips-to-make-a-or-b-equal-to-c/

//Greedy 贪心
//由于位运算是各位独立的运算，位于位之间不会有相互干扰（如进位）。我们按位考虑这个问题，把每个数位对答案的贡献计算清楚，然后累加即可。
//
//a=0,b=0,c=0 ans+=0
//a=0,b=0,c=1 ans+=1
//a=0,b=1,c=0 ans+=1
//a=0,b=1,c=1 ans+=0
//a=1,b=0,c=0 ans+=1
//a=1,b=0,c=1 ans+=0
//a=1,b=1,c=0 ans+=2
//a=1,b=1,c=1 ans+=0
//此外还需要最基本的位运算知识：
//
//一个数 xx 的二进制表示中第 ii 位为 00 当且仅当 x&(2^i)==0 (&为按位与)
//一个数 xx 的二进制表示中第 ii 位为 11 当且仅当 x&(2^i)!=0 (&为按位与)

//参考解答：https://leetcode-cn.com/circle/discuss/FagUOB/

//func minFlips(a int, b int, c int) int {
//	var size uint = 63
//	var i uint = 0
//	a_uint := uint(a)
//	b_uint := uint(b)
//	c_uint := uint(c)
//	result := 0
//	for ; i<size; i++ {
//		if ((1<<i) & c_uint) != 0 {
//			if ((1<<i) & a_uint) == 0 && ((1<<i) & b_uint) == 0 {result++}
//		} else {
//			if ((1<<i) & a_uint) != 0 {result++}
//			if ((1<<i) & b_uint) != 0 {result++}
//		}
//	}
//	return result
//}

//根据c的当前位，分两种情况讨论
//
//当前位位1，则a和b至少一个为1，如果两个都为0则修改任一个即可，修改位+1
//当前位为0，则a和b都必须为0，任一不为0修改位+1
//按位统计直到所有数都为0

func minFlips(a int, b int, c int) int {
	Ans := 0
	for a > 0 || b > 0 || c > 0 {
		if c&1 == 1 {
			if (a&1 == 0) && (b&1 == 0) {
				Ans = Ans + 1
			}
		}
		if c&1 == 0 {
			if a&1 == 1 {
				Ans = Ans + 1
			}
			if b&1 == 1 {
				Ans = Ans + 1
			}
		}
		a = a / 2
		b = b / 2
		c = c / 2
	}
	return Ans
}

func main() {
	fmt.Println(minFlips(2, 6, 5))
	fmt.Println(minFlips(4, 2, 7))
	fmt.Println(minFlips(1, 2, 3))
}
