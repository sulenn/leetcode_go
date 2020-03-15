package main

//https://leetcode-cn.com/problems/three-steps-problem-lcci/

func waysToStep(n int) int {
	if n == 1 {return 1}
	if n == 2 {return 2}
	if n == 3 {return 4}
	f1 := 1
	f2 := 2
	f3 := 4
	for i:=3; i<n; i++ {
		f1,f2,f3 = f2, f3, f1+f2+f3
		if f3 > 1000000007 {
			f3 = f3 % 1000000007
		}
	}
	return f3
}