package main

//https://leetcode-cn.com/problems/can-place-flowers/

func canPlaceFlowers(flowerbed []int, n int) bool {
	start, end := -1, -1
	for i := 0; i < len(flowerbed); i++ {
		if n <= 0 {
			return true
		}
		if flowerbed[i] == 0 {
			end = i
		}
		if flowerbed[i] == 1 {
			start, end = i+1, i+1
		}
		if end-start+1 == 3 {
			start = i
			n--
		}
	}
	if end-start+1 == 2 {
		n--
	}
	if n <= 0 {
		return true
	}
	return false
}
