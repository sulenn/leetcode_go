package main

//https://leetcode-cn.com/problems/rotate-matrix-lcci/

func rotate(matrix [][]int)  {
	n := len(matrix)
	//for i:=0; i<n/2; i++ {   // 原地反转，重要是知道转移方程式
	//	for j:=i; j<n-i-1; j++ {
	//		swap(&matrix[i][j], &matrix[j][n-i-1])
	//		swap(&matrix[i][j], &matrix[n-i-1][n-j-1])
	//		swap(&matrix[i][j], &matrix[n-j-1][i])
	//	}
	//}
	arr := make([][]int,n)   // 借助第二个数组
	for i:=0;i<n;i++ {
		arr[i] = make([]int,n)
	}
	for i:=0;i<n;i++ {
		for j:=0;j<n;j++ {
			arr[j][n-i-1] = matrix[i][j]
		}
	}
	for i:=0; i<n;i++ {
		copy(matrix[i],arr[i])
	}
}
