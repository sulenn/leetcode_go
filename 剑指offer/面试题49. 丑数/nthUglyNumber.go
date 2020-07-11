package main

//https://leetcode-cn.com/problems/chou-shu-lcof/

func nthUglyNumber(n int) int {
	if n < 1 {
		return 0
	}
	result := []int{1}
	l1, l2, l3 := 0, 0, 0
	for n > 1 {
		if result[l1]*2 <= result[l2]*3 && result[l1]*2 <= result[l3]*5 {
			if result[l1]*2 > result[len(result)-1] {
				result = append(result, result[l1]*2)
				n--
			}
			l1++
			continue
		}
		if result[l2]*3 <= result[l1]*2 && result[l2]*3 <= result[l3]*5 {
			if result[l2]*3 > result[len(result)-1] {
				result = append(result, result[l2]*3)
				n--
			}
			l2++
			continue
		} else {
			if result[l3]*5 > result[len(result)-1] {
				result = append(result, result[l3]*5)
				n--
			}
			l3++
		}
	}
	return result[len(result)-1]
}
