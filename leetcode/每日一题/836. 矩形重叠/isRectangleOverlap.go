package main

//https://leetcode-cn.com/problems/rectangle-overlap/

//反证法，先处理不重叠的四种情况
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	if len(rec1) != 4 || len(rec2) != 4 {return false}
	if rec1[2] <= rec2[0] || rec1[0] >= rec2[2] || rec1[1] >= rec2[3] || rec1[3] <= rec2[1] {return false}   // 不重叠的四种情况
	return true
}
