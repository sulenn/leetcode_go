package main

//https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/

// dp
//如果第i个字符之前没有出现过，那么 f(i)=f(i-1)+1。
//如果第 i 个字符之前出现过，同时这两个字符之间的距离 d 小于或等于 f(i-1)，那么 f(i)=d
//如果第 i 个字符之前出现过，同时这两个字符之间的距离 d 大于 f(i-1)，那么 f(i)=f(i-1)+1
//详细解释说明可以参考剑指 Offer
func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	curMaxLength := 0
	alphaDic := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		index, ok := alphaDic[s[i]]
		if !ok {
			curMaxLength++
		} else if i-index > curMaxLength {
			curMaxLength++
		} else if i-index <= curMaxLength {
			curMaxLength = i - index
		}
		alphaDic[s[i]] = i
		if curMaxLength > maxLength {
			maxLength = curMaxLength
		}
	}
	return maxLength
}
