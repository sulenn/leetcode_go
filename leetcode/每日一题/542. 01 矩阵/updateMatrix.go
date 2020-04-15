package main

//https://leetcode-cn.com/problems/01-matrix/

//广搜，参考：https://leetcode-cn.com/problems/01-matrix/solution/01ju-zhen-by-wishbuaa/
//也可以用两次dp：https://leetcode-cn.com/problems/01-matrix/solution/jian-dan-de-liang-ci-bian-li-jie-jue-by-habbi/
func updateMatrix(matrix [][]int) [][]int {
	arr := make([][]int, 0)
	flag := make([][]bool, len(matrix))
	for i:=0; i<len(matrix); i++ {   // 标志位，记录已访问的节点
		flag[i] = make([]bool, len(matrix[0]))
	}
	for i:=0; i<len(matrix); i++ {   // 记录全0的坐标
		for j:=0; j<len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				arr = append(arr, []int{i, j})
				flag[i][j] = true
			}
		}
	}
	location := [][]int {{1,0}, {-1,0}, {0,-1}, {0,1}}
	count := 0
	for len(arr) != 0 {
		nextLevel := make([][]int, 0)
		for i:=0; i<len(arr); i++ {  // 记录下一个层级的坐标
			matrix[arr[i][0]][arr[i][1]] = count
			for _, v := range location {
				x := arr[i][0]+v[0]
				y := arr[i][1]+v[1]
				if x >=0 && x < len(matrix) && y>=0 && y<len(matrix[0]) && !flag[x][y] {
					nextLevel = append(nextLevel, []int {x,y})
					flag[x][y] = true
				}
			}
		}
		arr = nextLevel
		count++
	}
	return matrix
}

