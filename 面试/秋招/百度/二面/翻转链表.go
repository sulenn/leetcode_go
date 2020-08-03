package main

func reverse(list *Node) *Node {
	if list == nil || list.Next == nil {
		return list
	}
	curNode := list
	nextNode := list.Next
	curNode.Next = nil
	for nextNode != nil {
		temp := nextNode.Next
		nextNode.Next = curNode
		curNode, nextNode = nextNode, temp
	}
	return curNode
}
