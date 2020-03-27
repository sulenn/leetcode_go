package main

//https://leetcode-cn.com/problems/can-place-flowers/

func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerbed = append([]int{0}, append(flowerbed, 0)...)
	for i:=1; i<len(flowerbed)-1 && n > 0; i++ {
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
			i++
		}
	}
	return n == 0
}
