package main

// hash map 解法
// 另外一种动态规划解法，参考：https://leetcode-cn.com/problems/longest-consecutive-sequence/solution/dong-tai-gui-hua-python-ti-jie-by-jalan/
func longestConsecutive(nums []int) int {
	numDic := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		numDic[nums[i]] = true
	}
	max := 0
	for k, v := range numDic {
		if !v {
			continue
		}
		len := 1
		num := k
		numDic[num] = false
		for true { // 减减减
			num--
			if v, ok := numDic[num]; ok && v {
				len++
				numDic[num] = false
			} else {
				break
			}
		}
		num = k
		for true { // 加加加
			num++
			if v, ok := numDic[num]; ok && v {
				len++
				numDic[num] = false
			} else {
				break
			}
		}
		if len > max {
			max = len
		}
	}
	return max
}
