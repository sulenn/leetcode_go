package main

type Node struct {  //节点结构体
	Value int
	Next []*Node
}

// 深度搜索的方法解决该问题
func longest(root *Node) int {
	if root == nil {return 0}
	result := 1   // 根节点也算一次子序列
	var dfs func(root *Node, counter int, pre int)  // 函数声名
	dfs = func(root *Node, counter int, pre int) {  // 深搜函数实现
		if result < counter {result = counter}  // 每次判断最大子序列的值
		if root == nil {return}
		if root.Value == pre+1 {    // 子节点的值比父节点值=+1
			for i:=0; i< len(root.Next); i++ {   // 广搜
				dfs(root.Next[i], counter+1,root.Value)
			}
		} else {
			for i:=0; i< len(root.Next); i++ {   // 从 1 开始计时
				dfs(root.Next[i], 1,root.Value)
			}
		}
	}
	for i:=0; i< len(root.Next); i++ {
		dfs(root.Next[i], 1, root.Value)
	}
	return result
}

