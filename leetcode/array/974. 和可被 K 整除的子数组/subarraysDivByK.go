package main

//https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/

//数组前缀和，参考：https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/solution/he-ke-bei-kzheng-chu-de-zi-shu-zu-by-lenn123/
func subarraysDivByK(A []int, K int) int {
	sum := 0
	dic := make(map[int]int)
	dic[0] = 1 // 字典初始化
	count := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
		sum = (sum%K + K) % K // 注意这儿需要考虑负数，所以 + K
		if _, ok := dic[sum]; ok {
			count += dic[sum]
		}
		dic[sum]++
	}
	return count
}
