package main

//https://leetcode-cn.com/problems/largest-1-bordered-square/

// 这个博主的很容易理解（解法不同）：https://leetcode-cn.com/problems/largest-1-bordered-square/solution/golang-20ms-jian-dan-yi-dong-by-resara/
// 另外一个 dp 解法：https://leetcode-cn.com/problems/largest-1-bordered-square/solution/dpzui-jian-dan-de-dai-ma-by-ri-mu-tu-yuan-12/
func largest1BorderedSquare(grid [][]int) int {
	r := len(grid)
	c := len(grid[0])
	max := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] != 1 {
				continue
			}
			length := 1
			for i+length < r && j+length < c { // 获取当前点开始的最长边长，边长满足水平和垂直均为1
				if grid[i+length][j] == 1 && grid[i][j+length] == 1 {
					length++
				} else {
					break
				}
			}
			length-- // 防止溢出
			for length > 0 {
				flag := true
				for v := 1; v <= length; v++ { // 满足正方型条件的最长边长
					if grid[i+v][j+length] != 1 || grid[i+length][j+v] != 1 {
						flag = false
						break
					}
				}
				if flag {
					break
				}
				length--
			}
			if max < (length+1)*(length+1) {
				max = (length + 1) * (length + 1)
			} // 更新最大正方形面积
		}
	}
	return max
}
