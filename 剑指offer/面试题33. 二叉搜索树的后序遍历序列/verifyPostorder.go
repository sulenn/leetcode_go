package main

//https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/

//最后一个节点为根节点，左子树节点值均小于根节点，右子树节点均大于根节点值
func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 || len(postorder) == 1{return true}
	leftTreeRage := 0
	for ;leftTreeRage < len(postorder)-1;leftTreeRage++ {   // 获取左子树范围
		if postorder[leftTreeRage] > postorder[len(postorder)-1] {
			break
		}
	}
	rightTreeRage := leftTreeRage
	for ;rightTreeRage < len(postorder)-1; rightTreeRage++ {   // 获取右子树范围
		if postorder[rightTreeRage] < postorder[len(postorder)-1] {
			return false
		}
	}
	return verifyPostorder(postorder[:leftTreeRage]) && verifyPostorder(postorder[leftTreeRage:rightTreeRage])
}