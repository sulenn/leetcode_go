package main

//https://leetcode-cn.com/problems/distribute-candies-to-people/

func distributeCandies(candies int, num_people int) []int {
	result := make([]int, num_people)
	countPer := 1 // 每次分糖果的数量
	count := 0    // 控制循环
	for candies > 0 {
		if candies < countPer {
			countPer = candies
		}
		result[count%(num_people)] += countPer
		count++
		candies -= countPer
		countPer++
	}
	return result
}
