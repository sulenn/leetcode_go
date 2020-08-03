package main

type Node struct {
	Value int
	Next  *Node
}

func listAdd(list1 *Node, list2 *Node) *Node {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	curNode := &Node{0, nil}
	head := curNode
	rest := 0
	for list1 != nil && list2 != nil {
		rest = (list1.Value + list2.Value + rest) / 10
		value := (list1.Value + list2.Value + rest) % 10
		temp := &Node{value, nil}
		curNode.Next = temp
		curNode = temp
		list1 = list1.Next
		list2 = list2.Next
	}
	for ; list1 != nil; list1 = list1.Next {
		rest = (list1.Value + rest) / 10
		value := (list1.Value + rest) % 10
		temp := &Node{value, nil}
		curNode.Next = temp
		curNode = temp
	}
	for ; list2 != nil; list2 = list2.Next {
		rest = (list2.Value + rest) / 10
		value := (list2.Value + rest) % 10
		temp := &Node{value, nil}
		curNode.Next = temp
		curNode = temp
	}
	if rest > 0 {
		temp := &Node{rest, nil}
		curNode.Next = temp
	}
	return head.Next
}
