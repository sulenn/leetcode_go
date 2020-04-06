package main

import "debug/macho"

type node struct {
	val int
	left *node
	right *node
}

func judgement(root *node) bool {

}

func judge(root *node, min int, max int) bool {
	if root == nil {return true}
	return judge(root.left, min, root.val) && judge(root.right, root.val, max)
}