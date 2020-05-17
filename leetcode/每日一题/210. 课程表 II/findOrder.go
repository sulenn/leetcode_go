package main

//https://leetcode-cn.com/problems/course-schedule-ii/

func findOrder(numCourses int, prerequisites [][]int) []int {
	node := make([]int, numCourses)
	for i:=0; i<len(prerequisites); i++ {
		for j:=len(prerequisites[i])-2; j>=0; j-- {  // 记录节点入度情况
			node[prerequisites[i][j]]++
		}
	}
	learningSort := make([]int,0)  // 学习顺序
	for len(learningSort) != numCourses {
		preLen := len(learningSort)
		for i:=0; i<numCourses; i++ {
			if node[i] == 0 {
				learningSort = append(learningSort, i)
				node[i] = -1
				break
			}
		}
		if preLen == len(learningSort) {return []int{}}  // 当前轮次没有入度为0
		for i:=0; i<len(prerequisites); i++ {
			length := len(prerequisites[i])
			if prerequisites[i][length-1] == learningSort[len(learningSort)-1] && length > 1 {
				node[prerequisites[i][length-2]]--   // 节点入度减 1
				prerequisites[i] = prerequisites[i][:length-1]
			}
		}
	}
	if len(learningSort) == numCourses {return learningSort}
	return []int {}
}
