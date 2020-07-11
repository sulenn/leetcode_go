package main

import "fmt"

//https://leetcode-cn.com/contest/weekly-contest-171/problems/number-of-operations-to-make-network-connected/

//超出时间限制
//func makeConnected(n int, connections [][]int) int {
//	if len(connections) < n - 1 {   // 线数量少于点数量减一
//		return -1
//	}
//	computeList := make([]int, n)    // 用于判断某个点是否被访问
//	computeSum := 0        // 用于循环结束的条件，是否每个电脑都访问
//	result := 0    // 图中连通图数量
//	for computeSum < n {    // 电脑没有访问完
//		graphNode := []int {}   // 当前连通图中存在的电脑
//		for i, value := range computeList {   // 未被访问的电脑中序号最小的数
//			if value == 0 {
//				graphNode = append(graphNode, i)
//				computeList[i] = 1  // 标记电脑被访问
//				computeSum++
//				break;
//			}
//		}
//		for len(graphNode) != 0 {   // 当前连通图中存在未被遍历的电脑
//			node := graphNode[0]
//			graphNode = graphNode[1:]   // 移除当前电脑
//			for _, connection := range connections {    // 当前电脑序号在总的连线中出现的情况
//				if connection[0] == node && computeList[connection[1]] == 0 {   // 连线第一个结点为当前电脑序号，且第二个节点未被访问过
//					graphNode = append(graphNode, connection[1])
//					computeList[connection[1]] = 1
//					computeSum++
//				} else if connection[1] == node && computeList[connection[0]] == 0 {    // 连线第二个结点为当前电脑序号，且第一个节点未被访问过
//					graphNode = append(graphNode, connection[0])
//					computeList[connection[0]] = 1
//					computeSum++
//				}
//			}
//		}
//		result++
//	}
//	return result - 1  // 需要移动的线的数量等于连通图数量减一
//}

func main() {
	fmt.Println(makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}))
	fmt.Println(makeConnected(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}))
	fmt.Println(makeConnected(5, [][]int{{0, 1}, {0, 2}, {3, 4}, {2, 3}}))
	fmt.Println(makeConnected(100, [][]int{{17, 51}, {33, 83}, {53, 62}, {25, 34}, {35, 90}, {29, 41}, {14, 53}, {40, 84}, {41, 64}, {13, 68}, {44, 85}, {57, 58}, {50, 74}, {20, 69}, {15, 62}, {25, 88}, {4, 56}, {37, 39}, {30, 62}, {69, 79}, {33, 85}, {24, 83}, {35, 77}, {2, 73}, {6, 28}, {46, 98}, {11, 82}, {29, 72}, {67, 71}, {12, 49}, {42, 56}, {56, 65}, {40, 70}, {24, 64}, {29, 51}, {20, 27}, {45, 88}, {58, 92}, {60, 99}, {33, 46}, {19, 69}, {33, 89}, {54, 82}, {16, 50}, {35, 73}, {19, 45}, {19, 72}, {1, 79}, {27, 80}, {22, 41}, {52, 61}, {50, 85}, {27, 45}, {4, 84}, {11, 96}, {0, 99}, {29, 94}, {9, 19}, {66, 99}, {20, 39}, {16, 85}, {12, 27}, {16, 67}, {61, 80}, {67, 83}, {16, 17}, {24, 27}, {16, 25}, {41, 79}, {51, 95}, {46, 47}, {27, 51}, {31, 44}, {0, 69}, {61, 63}, {33, 95}, {17, 88}, {70, 87}, {40, 42}, {21, 42}, {67, 77}, {33, 65}, {3, 25}, {39, 83}, {34, 40}, {15, 79}, {30, 90}, {58, 95}, {45, 56}, {37, 48}, {24, 91}, {31, 93}, {83, 90}, {17, 86}, {61, 65}, {15, 48}, {34, 56}, {12, 26}, {39, 98}, {1, 48}, {21, 76}, {72, 96}, {30, 69}, {46, 80}, {6, 29}, {29, 81}, {22, 77}, {85, 90}, {79, 83}, {6, 26}, {33, 57}, {3, 65}, {63, 84}, {77, 94}, {26, 90}, {64, 77}, {0, 3}, {27, 97}, {66, 89}, {18, 77}, {27, 43}}))
}
