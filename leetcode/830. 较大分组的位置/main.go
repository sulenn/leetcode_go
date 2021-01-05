package main

//https://leetcode-cn.com/problems/positions-of-large-groups/

func largeGroupPositions(s string) [][]int {
	var result [][]int
	start, end := 0, 0
	var alph uint8 = '0'
	for i := 0; i < len(s); i++ {
		if alph == s[i] {
			end = i
			continue
		}
		if end-start+1 >= 3 {
			result = append(result, []int{start, end})
		}
		start, end = i, i
		alph = s[i]
	}
	if end-start+1 >= 3 {
		result = append(result, []int{start, end})
	}
	return result
}
