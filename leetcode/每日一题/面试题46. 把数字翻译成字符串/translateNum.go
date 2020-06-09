package main

import "strconv"

func translateNum(num int) int {
	dp1, dp2 := 0,0
	numStr := strconv.Itoa(num)
	for i:=0; i<len(numStr); i++ {
		if dp2 == 0 {
			dp1, dp2 = 1, 1
			continue
		}
		curNUm, _ := strconv.Atoi(numStr[i-1:i+1])
		if curNUm >= 10 && curNUm <= 25 {
			dp2, dp1 = dp2+dp1, dp2
		} else {
			dp2, dp1 = dp2, dp2
		}
	}
	return dp2
}