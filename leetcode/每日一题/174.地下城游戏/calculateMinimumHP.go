package main

//https://leetcode-cn.com/problems/dungeon-game/

//参考：https://leetcode-cn.com/problems/dungeon-game/solution/jian-dan-dfskan-yi-yan-jiu-ming-bai-e-by-sweetiee/
func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 {
		return 1
	}
	r := len(dungeon)
	c := len(dungeon[0])
	memo := make([][]int, r)
	for i := 0; i < r; i++ {
		memo[i] = make([]int, c)
	}
	var dfs func(m int, n int) int
	dfs = func(m int, n int) int {
		if m == r-1 && n == c-1 {
			memo[m][n] = max(1-dungeon[m][n], 1)
		}
		if memo[m][n] > 0 {
			return memo[m][n]
		}
		var minvalue int
		if m == r-1 {
			minvalue = max(dfs(m, n+1)-dungeon[m][n], 1)
		} else if n == c-1 {
			minvalue = max(dfs(m+1, n)-dungeon[m][n], 1)
		} else {
			minvalue = min(dfs(m+1, n)-dungeon[m][n], dfs(m, n+1)-dungeon[m][n])
			minvalue = max(minvalue, 1)
		}
		memo[m][n] = minvalue
		return minvalue
	}
	dfs(0, 0)
	return memo[0][0]
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
