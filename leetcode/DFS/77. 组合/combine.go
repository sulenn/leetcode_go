package main

//https://leetcode-cn.com/problems/combinations/

//组合
func combine(n int, k int) [][]int {
	result := [][]int{}
	var dfs func(start int, arr []int)
	dfs = func(start int, arr []int) {
		if len(arr) == k {
			deepCopy := make([]int, len(arr))
			copy(deepCopy, arr) // 注意这儿要用深拷贝
			result = append(result, deepCopy)
			return
		}
		for i := start; i <= n; i++ {
			dfs(i+1, append(arr, i))
		}
	}
	dfs(1, []int{})
	return result
}
