package main

import "fmt"

//https://leetcode-cn.com/contest/weekly-contest-169/problems/jump-game-iii/

//![image.png](https://ws1.sinaimg.cn/large/006alGmrgy1gakrnaec2ij30wr0nuwg5.jpg)

//广搜
func canReach(arr []int, start int) bool {
	nodeList := []int{start}
	length := len(arr)
	for len(nodeList) > 0 {
		node := nodeList[0]
		nodeList = nodeList[1:]
		if arr[node] == 0 { // 满足条件
			return true
		}
		if arr[node] != -1 && node+arr[node] < length { // 排除值为 -1 节点，然后向前跳
			nodeList = append(nodeList, node+arr[node])
		}
		if arr[node] != -1 && node-arr[node] >= 0 { // 排除值为 -1 节点，然后向后跳
			nodeList = append(nodeList, node-arr[node])
		}
		arr[node] = -1 // 已经访问过的节点置为 -1
	}
	return false
}

func main() {
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5))
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 0))
	fmt.Println(canReach([]int{3, 0, 2, 1, 2}, 2))
}
