package main

import "sort"

//https://leetcode-cn.com/problems/assign-cookies/

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	result := 0
	for i,j :=0,0; i < len(g) && j <len(s); {
		if g[i] <= s[j] {
			result++
			i++
		}
		j++
	}
	return result
}
