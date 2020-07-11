package main

//https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/

func singleNumbers(nums []int) []int {
	first := 0 // 获取异或值
	for i := 0; i < len(nums); i++ {
		first ^= nums[i]
	}
	num := 1
	for first&num == 0 { // 获取异或值中从右往左的第一个 1 所处的位置
		num <<= 1
	}
	result := make([]int, 2)
	for _, v := range nums {
		if v&num == 0 { // 数组中与 num 异或结果为0 的分为一组
			result[0] ^= v
		} else { // 不为 0 的分为一组
			result[1] ^= v
		}
	}
	return result
}
