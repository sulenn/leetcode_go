package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	curNode := head
	for curNode != nil { // 将新节点塞入原节点后面
		newNode := &Node{Val: curNode.Val}
		newNode.Next = curNode.Next
		curNode.Next = newNode
		curNode = curNode.Next.Next
	}
	curNode = head
	for curNode != nil { // 拷贝新节点的 Random
		if curNode.Random != nil {
			curNode.Next.Random = curNode.Random.Next
		}
		curNode = curNode.Next.Next
	}
	newNode := head.Next
	curNode = newNode
	for head != nil { // 将新节点从链表中分离，注意旧链表中不能出现新节点，需要把新节点剥离
		head.Next = head.Next.Next
		head = head.Next
		if curNode.Next != nil {
			curNode.Next = curNode.Next.Next
		}
		curNode = curNode.Next
	}
	return newNode
}
