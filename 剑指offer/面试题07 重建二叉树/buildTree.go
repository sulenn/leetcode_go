package main

//https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/

import "errors"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历的一个元素是根节点
//根节点在中序遍历中所处位置的数组左边是左子树，右边是右子树。左右子树满足中序遍历
//通过左、右子树的节点数量可以从 preorder 中获得左、右子树的前序遍历
//满足递归条件
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		panic(errors.New("前序遍历和中序遍历的数量不一致！"))
	}
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{preorder[0], nil, nil}
	}
	root := &TreeNode{preorder[0], nil, nil}
	leftTreeNodeNums := 0           // 左子树节点的数量
	for i, value := range inorder { // 获取根节点在中序遍历数组中的位子
		if value == preorder[0] {
			leftTreeNodeNums = i
			break
		}
	}
	root.Left = buildTree(preorder[1:leftTreeNodeNums+1], inorder[:leftTreeNodeNums])
	root.Right = buildTree(preorder[leftTreeNodeNums+1:], inorder[leftTreeNodeNums+1:])
	return root
}
