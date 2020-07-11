package main

import (
	"fmt"
	"math"
)

//https://leetcode-cn.com/problems/network-delay-time/

//迪杰斯特拉解法，时间复杂度 O(N^2)
func networkDelayTime(times [][]int, N int, K int) int {
	graph := make([][]int, N+1) // 邻接矩阵， 用于存放各个点到各个点的距离
	for i := 0; i < N+1; i++ {
		graph[i] = make([]int, N+1)
	}
	for i := 0; i < N+1; i++ {
		for j := 0; j < N+1; j++ {
			graph[i][j] = -1
		}
	}
	for _, node := range times { // 遍历times填充邻接表
		graph[node[0]][node[1]] = node[2]
	}
	minDistance := make(map[int]int, N)   // 存放 K 到各个点的最短路径，最大的那个最短路径即为结果
	minDistance[K] = 0                    // K 到 K的距离为0
	curDistance := make(map[int]int, N-1) // 存放当前 K 到各个点的最短路径, 共 N-1 个
	for i := 1; i < N+1; i++ {            // 初始化 curDistance 为 K 到各个节点的距离
		if i != K { // 排除 K
			curDistance[i] = graph[K][i]
		}
	}
	for i := 0; i <= N-2; i++ { // 还剩下 N - 1 个点待处理
		minValue := math.MaxInt64
		minKey := 1
		for key, value := range curDistance { // 遍历所有节点，找到离K最近的节点
			if value != -1 && value < minValue {
				minValue = value
				minKey = key
			}
		}
		if minValue == math.MaxInt64 { // 剩下结点都不可到达
			break
		}
		minDistance[minKey] = minValue // 添加新的 minDistance 中最短距离的节点
		delete(curDistance, minKey)    // 删除 curDistance 中最新的最短距离节点
		for j := 1; j <= N; j++ {
			if graph[minKey][j] != -1 { // 存在当前最小值节点 minKey 到 j 节点的路径
				if _, ok := curDistance[j]; ok { // j 节点没有被当作最小路径点删除过
					if curDistance[j] == -1 { // 之前从 K 结点无法到达该J结点
						curDistance[j] = minValue + graph[minKey][j]
					} else {
						if curDistance[j] > minValue+graph[minKey][j] {
							curDistance[j] = minValue + graph[minKey][j]
						}
					}
				}
			}
		}
	}
	if len(minDistance) != N { // 存在没有访问过的结点
		return -1
	}
	max := math.MinInt64
	for _, value := range minDistance { // K 点到其它结点的最长路径
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	//fmt.Println(networkDelayTime([][]int {{2,1,1},{2,3,1},{3,4,1}}, 4, 2))
	fmt.Println(networkDelayTime([][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 2}}, 3, 1))
	//fmt.Println(networkDelayTime([][]int {{1,2,1}}, 2, 2))
}
