package main

//https://leetcode-cn.com/problems/stone-game/

//dp[i][j] 表示，i-j取值中第一、二个玩家可以取得的最高分数
//参考：https://leetcode-cn.com/problems/stone-game/solution/jie-jue-bo-yi-wen-ti-de-dong-tai-gui-hua-tong-yong/
func stoneGame(piles []int) bool {
	length := len(piles)
	arr := make([][][2]int, length)
	for i:=0; i<length; i++ {
		arr[i] = make([][2]int, length)
	}
	for i:=length-1; i>=0; i-- {  // 注意要倒着来，从最高点往下，类似最大回文子序列和
		for j:=i; j<length; j++ {
			if i == j {   // 对角线，说明只有一个元素
				arr[i][j] = [2]int{piles[i], 0}
			} else {   // dp[i][j][0] = max(arr[i+1][j][1] + piles[i], arr[i][j-1][1] + piles[j])
				if arr[i+1][j][1] + piles[i] > arr[i][j-1][1] + piles[j] {
					arr[i][j][0] = arr[i+1][j][1] + piles[i]
					arr[i][j][1] = arr[i+1][j][0]
				} else {
					arr[i][j][0] = arr[i][j-1][1] + piles[j]
					arr[i][j][1] = arr[i][j-1][0]
				}
			}
		}
	}
	return arr[0][length-1][0] > arr[0][length-1][1]
}