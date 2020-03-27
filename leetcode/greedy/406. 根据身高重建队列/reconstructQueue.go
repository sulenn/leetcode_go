package main

import "sort"

//https://leetcode-cn.com/problems/queue-reconstruction-by-height/

//思路：按前方身高大于等于排序，排序完之后进行类似插入排序的操作
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][1] < people[j][1]
	})
	for i:=1; i<len(people); i++ {
		count := 0  // 前方身高大于当前的人数
		for j:=0; j<i; j++ {
			if people[i][0]<=people[j][0] && people[i][1]==count {  // 满足条件，插入
				temp := people[i]
				move(people[j:i+1])   // 腾出待插入的位置
				people[j] = temp
				break
			} else if people[i][0]<=people[j][0] {  // 身高大于当前身高
				count++
			}
		}
	}
	return people
}

func move(nums [][]int)  {
	for i:=len(nums)-1; i>0; i-- {
		nums[i] = nums[i-1]
	}
}
