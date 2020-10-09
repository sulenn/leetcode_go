package main

import (
	"fmt"
	"strconv"
	"strings"
)

//层序遍历，参考：https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/solution/ceng-xu-bian-li-by-tinylife/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) Serialize(root *TreeNode) string {
	result := make([]string, 0)
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]
		if curNode != nil {
			queue = append(queue, curNode.Left, curNode.Right)
			result = append(result, strconv.Itoa(curNode.Val))
		} else {
			result = append(result, "null")
		}
	}
	return strings.Join(result, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) Deserialize(data string) *TreeNode {
	stringArr := strings.Split(data, ",")
	if len(data) <= 1 {
		return nil
	}
	value, err := strconv.Atoi(stringArr[0])
	if err != nil {
		fmt.Println("error when transfer string to int")
		return nil
	}
	root := &TreeNode{Val: value}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]
		if curNode == nil {
			continue
		}
		left := stringArr[i]
		right := stringArr[i+1]
		i += 2
		if left != "null" { // 左子树
			leftValue, err := strconv.Atoi(left)
			if err != nil {
				fmt.Println("error when transfer string to int")
				return nil
			}
			curNode.Left = &TreeNode{Val: leftValue}
			queue = append(queue, curNode.Left)
		} else {
			queue = append(queue, nil)
		}
		if right != "null" { // 右子树
			rightValue, err := strconv.Atoi(right)
			if err != nil {
				fmt.Println("error when transfer string to int")
				return nil
			}
			curNode.Right = &TreeNode{Val: rightValue}
			queue = append(queue, curNode.Right)
		} else {
			queue = append(queue, nil)
		}
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
