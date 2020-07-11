package main

//https://leetcode-cn.com/problems/generate-parentheses/

//深搜
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	var dfs func(l, r int, str string)
	result := []string{}
	dfs = func(l, r int, str string) {
		if l == 0 && r == 0 {
			result = append(result, str)
		}
		if l > 0 {
			dfs(l-1, r, str+"(")
		}
		if r > l {
			dfs(l, r-1, str+")")
		}
	}
	dfs(n, n, "")
	return result
}
