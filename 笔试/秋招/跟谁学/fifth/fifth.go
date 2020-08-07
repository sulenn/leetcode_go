package main

import (
	"fmt"
	"io"
)

//时间限制： 3000MS
//内存限制： 589824KB
//题目描述：
//给定某篮球运动员连续N次三分投篮是否命中的记录，求其中连续命中至少5球的次数。
//
//
//
//输入描述
//第一行读入一个整数N，表示共有N次三分投篮记录
//
//第二行读入N个整数，1为投中，0为投失
//
//输出描述
//输出一行，包含一个整数，表示连续命中至少5次事件发生的次数

func main() {
	var N int
	for {
		_, err := fmt.Scan(&N)
		if err == io.EOF {
			break
		}
		nums := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&nums[i])
		}
		result := 0
		count := 0
		flag := true
		for i := 0; i < N; i++ {
			if nums[i] == 1 {
				count++
			} else {
				count = 0
				flag = true
			}
			if flag && count >= 5 {
				result++
				flag = false
			}
		}
		fmt.Println(result)
	}
}
