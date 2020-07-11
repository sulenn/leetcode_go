package main

//https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}
	pushedStack := []int{} // 模拟入栈
	poppedStartPoint := 0
	for i := 0; i < len(pushed); i++ {
		if pushed[i] == popped[poppedStartPoint] {
			poppedStartPoint++
			for len(pushedStack) != 0 {
				if pushedStack[len(pushedStack)-1] == popped[poppedStartPoint] {
					pushedStack = pushedStack[:len(pushedStack)-1]
					poppedStartPoint++
				} else {
					break
				}
			}
		} else {
			pushedStack = append(pushedStack, pushed[i])
		}
	}
	for len(pushedStack) != 0 {
		if pushedStack[len(pushedStack)-1] == popped[poppedStartPoint] {
			pushedStack = pushedStack[:len(pushedStack)-1]
			poppedStartPoint++
		} else {
			return false
		}
	}
	return true
}
