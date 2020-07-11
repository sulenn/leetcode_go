package main

//https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof/

//思路：当前回合骰子的点数N=上一回合骰子的点数 (N-5) + (N-4) + (N-3) + (N-2) + (N-1) + (N-0)
func twoSum(n int) []float64 {
	result := make([][]int, 2) // 0 为最新的回合点数
	result[0] = []int{1, 1, 1, 1, 1, 1}
	for i := 2; i <= n; i++ {
		result[1] = make([]int, i*6-i+1)
		for j := 0; j < i*6-i+1; j++ {
			if j-0 >= 0 && j-0 < len(result[0]) {
				result[1][j] += result[0][j-0]
			}
			if j-1 >= 0 && j-1 < len(result[0]) {
				result[1][j] += result[0][j-1]
			}
			if j-2 >= 0 && j-2 < len(result[0]) {
				result[1][j] += result[0][j-2]
			}
			if j-3 >= 0 && j-3 < len(result[0]) {
				result[1][j] += result[0][j-3]
			}
			if j-4 >= 0 && j-4 < len(result[0]) {
				result[1][j] += result[0][j-4]
			}
			if j-5 >= 0 && j-5 < len(result[0]) {
				result[1][j] += result[0][j-5]
			}
		}
		result[0] = result[1]
	}
	sum := 0
	for i := 0; i < len(result[0]); i++ { // 当前回合各点数的总量
		sum += result[0][i]
	}
	finalResult := make([]float64, len(result[0]))
	for i := 0; i < len(result[0]); i++ {
		finalResult[i] = float64(result[0][i]) / float64(sum)
	}
	return finalResult
}
