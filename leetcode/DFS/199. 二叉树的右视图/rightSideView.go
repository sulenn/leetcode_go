package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	return dfs(root, 0, []int {})
}

// 深搜的顺序为根节点、右子树、左子树
func dfs(root *TreeNode, depth int, rightSideViewValue []int) []int {
	if root == nil {
		return []int {}
	}
	depth++
	if len(rightSideViewValue) < depth {
		rightSideViewValue = append(rightSideViewValue, root.Val)
	}
	if root.Right != nil {
		rightSideViewValue = dfs(root.Right, depth, rightSideViewValue)
	}
	if root.Left != nil {
		rightSideViewValue = dfs(root.Left, depth, rightSideViewValue)
	}
	return rightSideViewValue
}

//func test(test [][]int)  {
//	//test = append(test,[]int {5})
//	test[0][0] = 9
//	fmt.Println(test)
//}
//
//func main()  {
//	temp := [][]int {{6}}
//	fmt.Println(temp)
//	test(temp)
//	fmt.Println(temp)
//}