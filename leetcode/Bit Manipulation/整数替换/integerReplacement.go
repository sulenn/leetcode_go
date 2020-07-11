package main

//https://leetcode-cn.com/problems/integer-replacement/

//偶数直接右移，只有一种选项
//奇数+1或者-1，有两种选项。
//2.1 显然，让每一步1的数目最少好处大，于是 0bxxx01 采用 -1； 0bxxx11 采用 +1；
//2.2 特殊情况 3，按上述原则+1后两次右移共需3次；减一后只需一次右移共2次，因此3采用-1操作

//参考：https://leetcode-cn.com/problems/integer-replacement/solution/wei-yun-suan-by-ma-xing/

func integerReplacement(n int) int {
	result := 0
	for n != 1 {
		if n&1 == 0 {
			n = n >> 1
		} else if n&2 == 0 || n == 3 {
			n--
		} else {
			n++
		}
		result++
	}
	return result
}
