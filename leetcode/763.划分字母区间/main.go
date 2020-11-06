package main

//https://leetcode-cn.com/problems/partition-labels/

//参考：https://leetcode-cn.com/problems/partition-labels/solution/shou-hua-tu-jie-hua-fen-zi-mu-qu-jian-ji-lu-zui-yu/
func partitionLabels(S string) []int {
	maxPositionArr := make([]int, 26)
	for i := 0; i < len(S); i++ {
		maxPositionArr[S[i]-'a'] = i
	}
	start := 0
	maxPosition := 0
	result := make([]int, 0)
	for i := 0; i < len(S); i++ {
		if maxPositionArr[S[i]-'a'] > maxPosition {
			maxPosition = maxPositionArr[S[i]-'a']
		}
		if i == maxPosition {
			result = append(result, maxPosition-start+1)
			start = maxPosition + 1
		}
	}
	return result
}
