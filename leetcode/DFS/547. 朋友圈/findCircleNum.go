package main

//https://leetcode-cn.com/problems/friend-circles/

// 深搜
func findCircleNum(M [][]int) int {
	visited := make([]bool, len(M))
	count := 0
	for i := 0; i < len(M); i++ {
		if !visited[i] {
			count++
			visited[i] = true
			dfs(M, visited, i)
		}
	}
	return count
}

func dfs(M [][]int, visited []bool, num int) {
	for i := 0; i < len(M); i++ {
		if M[num][i] == 1 && !visited[i] {
			visited[i] = true
			dfs(M, visited, i)
		}
	}
}
