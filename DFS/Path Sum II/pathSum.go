package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil && sum - root.Val == 0 {
		return [][]int {{root.Val}}
	}
	result := [][]int{}
	if root.Left != nil{
		left := pathSum(root.Left, sum - root.Val)   // 左子树递归
		if len(left) != 0 {
			for k ,v := range left {
				temp := root.Val   // 路径中添加根节点值
				for i := 0; i < len(v); i++ {
					temp1 := v[i]
					v[i] = temp
					temp = temp1
				}
				v = append(v, temp)
				result = append(result, v)
				k++
			}
		}
	}

	if root.Right != nil{
		right := pathSum(root.Right, sum - root.Val)
		if len(right) != 0 {
			for k ,v := range right {
				temp := root.Val
				for i := 0; i < len(v); i++ {
					temp1 := v[i]
					v[i] = temp
					temp = temp1
				}
				v = append(v, temp)
				result = append(result, v)
				k++
			}
		}
	}
	return result
}

func main() {
	
}
