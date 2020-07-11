package main

//https://leetcode-cn.com/problems/permutation-in-string/

// 判断 s 中是否存在 t 的排列
//参考：https://leetcode-cn.com/problems/permutation-in-string/solution/wo-xie-liao-yi-shou-shi-ba-suo-you-hua-dong-chuang/
func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		if sum, ok := need[s2[right]]; ok { // 进行窗口内数据的一系列更新
			window[s2[right]]++
			if window[s2[right]] == sum {
				valid++
			}
		}
		right++
		for right-left >= len(s1) { // 判断左侧窗口是否要收缩
			if valid == len(need) {
				return true
			} // 在这里判断是否找到了合法的子串
			if sum, ok := need[s2[left]]; ok { // 进行窗口内数据的一系列更新
				if window[s2[left]] == sum {
					valid--
				}
				window[s2[left]]--
			}
			left++
		}
	}
	// 未找到符合条件的子串
	return false
}
