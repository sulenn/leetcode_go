package main

func GetCoinCount(N int) int {
	// write code here
	result := 0
	money := 1024 - N
	result += money / 64
	money = money % 64
	result += money / 16
	money = money % 16
	result += money / 4
	money = money % 4
	result += money / 1
	return result
}
