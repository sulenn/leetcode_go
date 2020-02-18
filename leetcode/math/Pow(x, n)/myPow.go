package main

//https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/

func myPow(x float64, n int) float64 {
	if isZero(x) && n < 0 {  // 异常，对 0 求倒数
		return 0.0
	}
	if n >= 0 {   // 0 的 0 次方是没有意义的，因此无论是输出 0 还是 1 都是可以接收的。
		return uintPow(x, uint(n))
	} else {
		return uintPow(1.0/x, uint(-n))
	}
}

//由于计算机表示小数都有误差，我们不能直接用等号（==）判断两个小数是否相等。
//如果两个小数差的绝对值很小，比如小于0.0000001，就可以人为它们相等。
func isZero(x float64) bool {
	if x - 0.0 > -0.0000001 && x - 0.0 < 0.0000001 {
		return true
	} else {
		return false
	}
}

//func uintPow(x float64, n uint) float64 {   // 正常循环求无符号整数的次方
//	result := 1.0
//	var i uint
//	for i = 0; i<n;i++ {
//		result *= x
//	}
//	return result
//}


func uintPow(x float64, n uint) float64 {   // 移位递归求无符号整数的次方
	if n == 0 {return 1.0}
	if n == 1 {return x}
	result := uintPow(x, n >> 1)
	result *= result
	if n & 1 == 1 {result *= x}
	return result
}