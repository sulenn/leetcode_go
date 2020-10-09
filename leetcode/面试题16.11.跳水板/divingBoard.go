package main

//https://leetcode-cn.com/problems/diving-board-lcci/

//数学题，参考：https://leetcode-cn.com/problems/diving-board-lcci/solution/qing-xi-yi-dong-de-shu-xue-tui-dao-che-di-nong-don/
func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 { // 特例
		return []int{}
	}
	if shorter == longer { // 特例
		return []int{shorter * k}
	}
	result := []int{}
	for i := 0; i <= k; i++ {
		value := shorter*k + (longer-shorter)*i
		result = append(result, value)
	}
	return result
}
