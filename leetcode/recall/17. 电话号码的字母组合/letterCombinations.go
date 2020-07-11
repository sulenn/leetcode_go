package main

//https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

func letterCombinations(digits string) []string {
	var dfs func(digits string, str string)
	dic := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	result := make([]string, 0)
	dfs = func(digits string, str string) {
		if len(digits) == 0 {
			if str != "" {
				result = append(result, str)
			}
			return
		}
		nums := dic[digits[0]]
		for i := 0; i < len(nums); i++ {
			dfs(digits[1:], str+string(nums[i]))
		}
	}
	dfs(digits, "")
	return result
}
