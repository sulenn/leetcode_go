package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	return backtracking(nums)
}

// 回溯
func backtracking(nums []int) []*TreeNode {
	if len(nums) == 1 {
		return []*TreeNode{&TreeNode{nums[0], nil, nil}}
	}
	curTrees := []*TreeNode{}
	for i := 0; i < len(nums); i++ {
		leftTrees := backtracking(nums[:i])
		rightTrees := backtracking(nums[i+1:])
		if len(leftTrees) == 0 { // 只存在右子树
			for j := 0; j < len(rightTrees); j++ {
				curTrees = append(curTrees, &TreeNode{nums[i], nil, rightTrees[j]})
			}
		} else if len(rightTrees) == 0 { // 只存在左子树
			for j := 0; j < len(leftTrees); j++ {
				curTrees = append(curTrees, &TreeNode{nums[i], leftTrees[j], nil})
			}
		} else { // 存在左子树和右子树
			for u := 0; u < len(leftTrees); u++ {
				for v := 0; v < len(rightTrees); v++ {
					curTrees = append(curTrees, &TreeNode{nums[i], leftTrees[u], rightTrees[v]})
				}
			}
		}
	}
	return curTrees
}
