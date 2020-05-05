package main

//https://leetcode-cn.com/problems/palindromic-substrings/


//从中心往两侧延伸
//在长度为 N 的字符串中，可能的回文串中心位置有 2N-1 个：字母，或两个字母中间。
//对于每个可能的回文串中心位置，尽可能扩大它的回文区间 [left, right]。
//当 left >= 0 and right < N and S[left] == S[right] 时，扩大区间。
//参考：https://leetcode-cn.com/problems/palindromic-substrings/solution/hui-wen-zi-chuan-by-leetcode/
//也可以用 dp 来做：https://leetcode-cn.com/problems/palindromic-substrings/solution/liang-dao-hui-wen-zi-chuan-de-jie-fa-xiang-jie-zho/
func countSubstrings(s string) int {
	length := len(s)
	count := 0
	for i:=0; i<2*length-1; i++ {
		left := i / 2   // left 和 right 有可能相等，也有可能 left+1 = right
		right := left + i%2
		for ;left >=0 && right < length && s[left] == s[right]; {
			count++
			left--   // 扩大可能的回文区间
			right++
		}
	}
	return count
}

