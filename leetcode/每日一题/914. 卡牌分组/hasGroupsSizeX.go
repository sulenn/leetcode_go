package main

//https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards/

func hasGroupsSizeX(deck []int) bool {
	if len(deck) == 0 {return false}
	dic := make(map[int]int)  // 字典，统计各数字频率
	for _,v := range deck {
		dic[v]++
	}
	num := 0
	for _, v := range dic {  // 用最大公约数来判断是否可以分组
		if num == 0 {
			num = v
			continue
		}
		if num != 1 && v % num == 0 {continue}
		num = maxRestNum(num, v)
		if num == 1 {return false}
	}
	return num != 1
}

// 求最大公约数
func maxRestNum(a,b int) int {
	for a - b != 0 {
		if a < b {
			a,b = b,a
		}
		a,b = a - b, b
	}
	return a
}