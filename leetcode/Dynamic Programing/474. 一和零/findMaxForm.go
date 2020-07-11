package main

//https://leetcode-cn.com/problems/ones-and-zeroes/description/

//二维01背包问题，背包容量为1和0的数量，可装入的物品为01组成的字符串
func findMaxForm(strs []string, m int, n int) int {
	arr := make([][]int, m+1) // 注意行长度为m+1，不然会出错。如第一个字符串为"10"
	for i := 0; i < m+1; i++ {
		arr[i] = make([]int, n+1) // 注意列长度为 n+1
	}
	for i := 0; i < len(strs); i++ {
		zeres := 0
		ones := 0
		for j := 0; j < len(strs[i]); j++ {
			if strs[i][j] == '1' {
				ones++
			} else {
				zeres++
			}
		}
		for u := m; u >= zeres; u-- {
			for v := n; v >= ones; v-- {
				if arr[u][v] < arr[u-zeres][v-ones]+1 {
					arr[u][v] = arr[u-zeres][v-ones] + 1
				}
			}
		}
	}
	return arr[m][n]
}

//思路，将字符串排序，从长度小的开始匹配。失败！！！比如["111","1000","1000","1000"]、9个1、3个0
//type Str struct {
//	s string
//	l int
//}

//func findMaxForm(strs []string, m int, n int) int {
//	str := []Str {}
//	for _, v := range strs {
//		str = append(str, Str{v, len(v)})
//	}
//	sort.Slice(str, func(i, j int) bool {
//		return str[i].l < str[j].l
//	})
//	result := 0
//	for _, v:= range str {
//		v_m, v_n := cal(v.s)
//		if m >= v_m && n >= v_n {
//			m -= v_m
//			n -= v_n
//			result++
//		}
//	}
//	return result
//}
//
//func cal(s string) (int, int) {
//	m, n :=0,0
//	for i:=0; i<len(s);i++ {
//		if s[i]=='0' {m++}
//		if s[i] == '1' {n++}
//	}
//	return m, n
//}
