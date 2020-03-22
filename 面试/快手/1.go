package main

import "fmt"

/**
 * 获取队中从前到后每个人与前方身高高于自己的人的最短距离
 * @param height int整型一维数组 队中从前到后每个人与前方身高高于自己的人的最短距离
 * @return int整型一维数组
 */
func DistanceToHigher( height []int ) []int {
	result := make([]int, len(height))
	p2 := len(height)-1
	for i:=len(height)-1; i>=0; i--  {
		p2 = i
		for p2 >=0 && height[p2] <= height[i] {
			p2--
		}
		if p2 == -1 {
			result[i] = 0
		} else {
			result[i] = i - p2
		}
	}
	return result
}

func main() {
	fmt.Println(DistanceToHigher([]int {}))
	fmt.Println(DistanceToHigher([]int {175,173,174,163,182,177}))
	fmt.Println(DistanceToHigher([]int {175, 179, 174, 163, 176, 177}))
}