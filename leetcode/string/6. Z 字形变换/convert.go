package main

//https://leetcode-cn.com/problems/zigzag-conversion/

func convert(s string, numRows int) string {
	if len(s) == 0 || numRows == 1 {
		return s
	} // 特殊情况
	zArr := make([][]string, numRows)
	zArr[0] = append(zArr[0], s[:1]) // zArr[0][0] 是规律之外的特例
	col := 0
	count := 1
	for count < len(s) { // 找规律
		if col%2 == 0 { // 奇数列的数量从 1 开始到 nums - 1
			for i := 2; i <= numRows; i++ {
				if count >= len(s) {
					break
				}
				zArr[i-1] = append(zArr[i-1], s[count:count+1])
				count++
			}
		} else { // 偶数列的数量从 nums-2 到 0
			for i := numRows - 1; i >= 1; i-- {
				if count >= len(s) {
					break
				}
				zArr[i-1] = append(zArr[i-1], s[count:count+1])
				count++
			}
		}
		col++
	}
	result := ""
	for i := 0; i < numRows; i++ { // 重新拼接 z 字形字符串
		for j := 0; j < len(zArr[i]); j++ {
			result += zArr[i][j]
		}
	}
	return result
}
