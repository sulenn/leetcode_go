package main

//https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum/

// 这道题目用双指针做比较好，参考：https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum/solution/zuo-you-zhi-zhen-wang-zhong-jian-bian-li-by-hlhfev/
// 不过我的这个方法也可用于四分、五分、六分......
func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, v := range A {
		sum += v
	}
	if sum % 3 != 0 || len(A) < 3 {return false}   // 数组长度大于3，且和能被 3 整除
	subSum := 0
	flag := 0
	for _, v := range A {
		subSum += v
		if subSum == sum / 3 {
			subSum = 0
			flag++
		}
	}
	if flag < 3 {return false}
	return true
}