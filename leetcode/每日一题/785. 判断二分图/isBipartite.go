package main

//https://leetcode-cn.com/problems/is-graph-bipartite/

//参考：https://leetcode-cn.com/problems/is-graph-bipartite/solution/shou-hua-tu-jie-bfs-si-lu-by-hyj8/
func isBipartite(graph [][]int) bool {
	visited := make(map[int]bool) // 用于记录访问过的顶点，同时记录染色情况，true为红色，false为蓝色
	queue := make([]int, 0)       //队列，用于广搜
	for i := 0; i < len(graph); i++ {
		if _, ok := visited[i]; ok { // 顶点是否已访问
			continue
		}
		queue = append(queue, i)
		visited[i] = true
		for len(queue) != 0 { // 广搜
			curNode := queue[0]
			curNodeColor := visited[curNode]
			queue = queue[1:]
			for j := 0; j < len(graph[curNode]); j++ {
				if color, ok := visited[graph[curNode][j]]; ok {
					if color != !curNodeColor {
						return false
					}
				} else {
					queue = append(queue, graph[curNode][j])
					visited[graph[curNode][j]] = !curNodeColor
				}
			}

		}
	}
	return true
}
