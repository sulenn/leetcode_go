package main

//https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//递归二叉树，字典累加。然后遍历字典。
//也可以用中序遍历的方法来做。注意左子树节点值小于等于根结点值，右子树节点值大于等于节点值，如（6,2,8,0,4,7,9,nil,nil,2,6）
func findMode(root *TreeNode) []int {
	result := []int {}
	node := make(map[int]int)
	var recursive func(root *TreeNode)
	recursive = func(root *TreeNode) {
		if root == nil {
			return
		}
		if _, ok:= node[root.Val]; ok {   // 字典累加
			node[root.Val]++
		} else {
				node[root.Val] = 1
		}
		recursive(root.Left)
		recursive(root.Right)
	}
	recursive(root)
	max := 0
	for key, value := range node {
		if value > max {
			result = []int {key}
			max = value
		} else if value == max {
			result = append(result, key)
		}
	}
	return result
}

func main() {
	//root := common.CreateTree([])
}
