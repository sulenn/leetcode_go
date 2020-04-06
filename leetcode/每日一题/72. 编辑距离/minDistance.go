package main

//https://leetcode-cn.com/problems/edit-distance/

//0/1背包问题
//状态定义：dp[i][j]表示word1的前i个字母转换成word2的前j个字母所使用的最少操作。
//状态转移：
//i指向word1，j指向word2
//若当前字母相同，则dp[i][j] = dp[i - 1][j - 1];
//否则取增删替三个操作的最小值 + 1， 即:
//dp[i][j] = min(dp[i][j - 1], dp[i - 1][j], dp[i - 1][j - 1]) + 1
//参考：https://leetcode-cn.com/problems/edit-distance/solution/jian-dan-dpmiao-dong-by-sweetiee/
func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {return len(word2)}
	if len(word2) == 0 {return len(word1)}
	arr := make([][]int, len(word1)+1)  // 初始化
	for i:=0; i<len(word1)+1;i++ {
		arr[i] = make([]int, len(word2)+1)
	}
	for i:=0; i<len(word1)+1; i++ {
		arr[i][0] = i
	}
	for j:=0; j<len(word2)+1; j++ {
		arr[0][j] = 0
	}
	for i:=1; i<len(word1)+1; i++ {
		for j:=1; j<len(word2)+1;j++ {
			if word1[i-1] == word2[j-1] {  // 状态转移
				arr[i][j] = arr[i-1][j-1]
			} else {
				arr[i][j] = MinValue(arr[i-1][j-1],arr[i-1][j],arr[i][j-1])+1
			}
		}
	}
	return arr[len(word1)][len(word2)]
}

func MinValue(v1, v2, v3 int) int {
	if v1 > v2 {
		v1 = v2
	}
	if v1 > v3 {
		v1 = v3
	}
	return v1
}