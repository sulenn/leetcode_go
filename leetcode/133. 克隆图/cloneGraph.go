package main

import "fmt"

//https://leetcode-cn.com/problems/clone-graph/

type Node struct {
	Val       int
	Neighbors []*Node
}

// 更好的思路，可以参考：https://leetcode-cn.com/problems/clone-graph/solution/go-bfsmap-by-bing-4/
// 更好的思路：DFS 递归，只需要一个 nodeMap 记录和标记克隆节点

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	queue := make([]*Node, 0)
	nodeMap := make(map[int]*Node) // 记录克隆的新节点
	flagMap := make(map[int]bool)  // 在第二个循环中标记已处理的节点
	queue = append(queue, node)
	for len(queue) > 0 { // 生成克隆节点和标记 map
		curNode := queue[0]
		queue = queue[1:]
		if _, ok := nodeMap[curNode.Val]; ok {
			continue
		}
		newNode := &Node{Val: curNode.Val, Neighbors: make([]*Node, len(curNode.Neighbors))}
		nodeMap[curNode.Val] = newNode
		flagMap[curNode.Val] = false
		queue = append(queue, curNode.Neighbors...)
	}
	queue = append(queue, node)
	for len(queue) > 0 { // 开始正式处理克隆
		curNode := queue[0]
		queue = queue[1:]
		if flagMap[curNode.Val] {
			continue
		}
		flagMap[curNode.Val] = true
		newNode := nodeMap[curNode.Val]
		for i := 0; i < len(curNode.Neighbors); i++ {
			queue = append(queue, curNode.Neighbors[i])
			newNode.Neighbors[i] = nodeMap[curNode.Neighbors[i].Val]
		}
	}
	return nodeMap[node.Val]
}

func main() {
	node0 := &Node{1, nil}
	//node1 := &Node{2, nil}
	//node2 := &Node{3, nil}
	//node3 := &Node{4, nil}
	//node0.Neighbors = []*Node{node1, node3}
	//node1.Neighbors = []*Node{node0, node2}
	//node2.Neighbors = []*Node{node1, node3}
	//node3.Neighbors = []*Node{node0, node2}
	node := cloneGraph(node0)
	fmt.Println("1")
	_ = node
}
