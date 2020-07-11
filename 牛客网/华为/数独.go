package main

func judge(arr [][]int) bool {
	if len(arr) != 9 || len(arr[0]) != 9 {
		return false
	}
	// 判断行
	for i := 0; i < 9; i++ {
		judgeArr := make([]bool, 9)
		for j := 0; j < 9; j++ {
			if arr[i][j] > 9 || arr[i][j] < 1 {
				return false
			}
			if judgeArr[arr[i][j]] {
				return false
			} else {
				judgeArr[arr[i][j]] = true
			}
		}
	}
	// 判断列
	for i := 0; i < 9; i++ {
		judgeArr := make([]bool, 9)
		for j := 0; j < 9; j++ {
			if judgeArr[arr[j][i]] {
				return false
			} else {
				judgeArr[arr[j][i]] = true
			}
		}
	}
	// 判断宫
	for i, j := 0, 0; i < 9 && j < 9; { // i 是行， j 是列
		judgeArr := make([]bool, 9)
		for u := 0; u < 3; u++ {
			for v := 0; v < 3; v++ {
				if judgeArr[arr[i+u][j+v]] {
					return false
				} else {
					judgeArr[arr[i+u][j+v]] = true
				}
			}
		}
		if j == 6 {
			j = 0
			i += 3
		} else {
			j += 3
		}
	}
	return true
}

func main() {
	//fmt.Println(judge())
}
