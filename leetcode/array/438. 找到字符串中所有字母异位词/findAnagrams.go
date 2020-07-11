package main

//https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/

func findAnagrams(s string, p string) []int {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	left, right := 0, 0
	valid := 0
	result := make([]int, 0)
	for right < len(s) {
		if sum, ok := need[s[right]]; ok { // 进行窗口内数据的一系列更新
			window[s[right]]++
			if window[s[right]] == sum {
				valid++
			}
		}
		right++
		for right-left >= len(p) { // 判断左侧窗口是否要收缩
			if valid == len(need) { // 在这里判断是否找到了合法的子串
				result = append(result, left)
			}
			if sum, ok := need[s[left]]; ok { // 进行窗口内数据的一系列更新
				if window[s[left]] == sum {
					valid--
				}
				window[s[left]]--
			}
			left++
		}
	}
	// 未找到符合条件的子串
	return result
}
