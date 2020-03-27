package main

//https://leetcode-cn.com/problems/partition-labels/description/

func partitionLabels(S string) []int {
	result := []int {}
	dic1 := make(map[byte]struct{})  // 独立区间中字符类型
	dic2 := make(map[byte]struct{})  // 独立区间外的字符类型
	p1 := 0  // p1 和 p2 是每个独立区间
	p2 := 0
	for ; p1 < len(S); {
		dic1 = make(map[byte]struct{})  // 初始化
		dic2 = make(map[byte]struct{})
		dic1[S[p1]] = struct{}{}
		for j:=p2+1; j < len(S); j++ {
			if _, ok := dic1[S[j]]; ok {   // 扩充原有区间
				for k,_ := range dic2 {
					dic1[k] = struct{}{}
				}
				p2 = j
			} else {
				dic2[S[j]] = struct{}{}
			}
		}
		result = append(result, p2-p1+1)
		p1 = p2+1
		p2 = p1
	}
	return result
}

