package Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal

import "fmt"

//* Definition for a binary tree node.
type TreeNode struct {
Val int
Left *TreeNode
Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		return nil
	}
	if len(preorder) == 0 {
		return nil
	}
	root := TreeNode{preorder[0], nil, nil}
	locate := 0
	//for i := 0; i < len(inorder); i++ {
	//	if preorder[0] == inorder[i] {
	//		locate = i
	//		break
	//	}
	//}
	for k, v := range inorder {
		if v == preorder[0] {
			locate = k
			break
		}
	}
	root.Left = buildTree(preorder[1:locate+1], inorder[0:locate])
	root.Right = buildTree(preorder[locate+1:], inorder[locate+1:])
	return &root
}

func main() {
	temp := buildTree([]int {3,9,20,15,7}, []int {9,3,15,20,7})
	fmt.Println(temp)
}
