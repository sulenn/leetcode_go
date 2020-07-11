package main

//https://leetcode-cn.com/problems/water-and-jug-problem/

//广搜，超时
//func canMeasureWater(x int, y int, z int) bool {
//	if x > x+y {return false}
//	if x == 0 || y == 0 {return y==z || x==z}
//	q :=make([][]int,0)  // 队列
//	visited := make([][]int, 0)   // 已访问过的状态
//	q = append(q, []int {0,0})
//	visited = append(visited,[]int{0,0})
//	for len(q) != 0 {
//		cur := q[0]
//		q = q[1:]
//		curX := cur[0]
//		curY := cur[1]
//		if curX == z || curY == z || curX+curY == z {return true}
//		// 桶一装满
//		if !judge([]int{x,curY}, visited) {
//			q = append(q, []int{x,curY})
//			visited = append(visited, []int{x,curY})
//		}
//		// 桶二装满
//		if !judge([]int{curX,y}, visited) {
//			q = append(q, []int{curX,y})
//			visited = append(visited, []int{curX,y})
//		}
//		// 桶一倒掉
//		if !judge([]int{0,curY}, visited) {
//			q = append(q, []int{0,curY})
//			visited = append(visited, []int{0,curY})
//		}
//		// 桶二装满
//		if !judge([]int{curX,0}, visited) {
//			q = append(q, []int{curX,0})
//			visited = append(visited, []int{curX,0})
//		}
//		//桶一向桶二倒
//		if curX+curY > y {
//			if !judge([]int{curX-y+curY, y}, visited) {
//				q = append(q, []int{curX-y+curY, y})
//				visited = append(visited, []int{curX-y+curY, y})
//			}
//		} else {
//			if !judge([]int{0, curX+curY}, visited) {
//				q = append(q, []int{0, curX+curY})
//				visited = append(visited, []int{0, curX+curY})
//			}
//		}
//	}
//	return false
//}

//func judge(state []int, visited [][]int) bool {
//	for _, v := range visited {
//		if state[0] == v[0] && state[1] == v[1] {
//			return true
//		}
//	}
//	return false
//}

//最大公约数
func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}
	if x == 0 || y == 0 {
		return x == z || y == z
	}
	if x > y {
		x, y = y, x
	}
	for y%x != 0 {
		y, x = x, y%x
	}
	return z%x == 0
}
