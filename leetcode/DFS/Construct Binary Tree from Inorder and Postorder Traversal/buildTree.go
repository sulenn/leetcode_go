package Construct_Binary_Tree_from_Inorder_and_Postorder_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) != len(postorder) {
		return nil
	}
	if len(inorder) == 0 {
		return nil
	}
	inorderLen := len(inorder)
	root := TreeNode{Val:postorder[inorderLen-1]}
	locate := 0
	for k, v := range inorder {
		if v == postorder[inorderLen-1] {
			locate = k
		}
	}
	root.Left = buildTree(inorder[:locate], postorder[:locate])
	root.Right = buildTree(inorder[locate+1:], postorder[locate:inorderLen-1])
	return &root
}

//func main() {
//
//}
