package main

//https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/

//宽度优先遍历
func movingCount(m int, n int, k int) int {
	arr := make([][]bool, m) // 数组，用于标记某个点是否已访问过
	for i := 0; i < m; i++ {
		arr[i] = make([]bool, n)
	}

	if m < 0 || n < 0 {
		return 0
	}
	if m == 0 && n == 0 {
		return 0
	}
	result := 0
	queue := [][]int{{0, 0}}
	for len(queue) != 0 {
		x := queue[0][0]
		y := queue[0][1]
		if sum(x)+sum(y) <= k && !arr[x][y] { // 注意已访问过的点不算
			result++
			if y+1 < n {
				queue = append(queue, []int{x, y + 1})
			}
			if x+1 < m {
				queue = append(queue, []int{x + 1, y})
			}
			arr[x][y] = true // 已访问
		}
		queue = queue[1:]
	}
	return result
}

func sum(num int) int {
	sum := 0
	for num != 0 {
		rest := num % 10
		sum += rest
		num = num / 10
	}
	return sum
}
