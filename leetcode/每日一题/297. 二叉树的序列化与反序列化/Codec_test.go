package main

import "testing"

func Test_Codec(t *testing.T)  {
	A := &TreeNode{Val: 2}
	B := &TreeNode{Val: 3,Left:&TreeNode{Val: 4},Right:&TreeNode{Val: 5}}
	root := &TreeNode{1,A,B}
	obj := Constructor()
	serializeStr := obj.Serialize(root)
	tree := obj.Deserialize(serializeStr)
	_ = tree
}