package main

//https://leetcode-cn.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/

func numSteps(s string) int {
	count := 0
	add := 0
	for s[0] == '0' {
		s = s[1:]
	}
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '0' && add == 0 {
			count += 1
		} else if s[i] == '0' && add == 1 {
			count += 2
		} else if s[i] == '1' && add == 1 {
			count += 1
		} else if s[i] == '1' && add == 0 {
			count += 2
			add = 1
		}
	}
	if add == 1 {
		count += 1
	}
	return count
}
