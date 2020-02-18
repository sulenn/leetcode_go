package main

import "fmt"

//https://leetcode-cn.com/problems/airplane-seat-assignment-probability/

//动态规划，参考 https://leetcode-cn.com/problems/airplane-seat-assignment-probability/solution/ju-ti-fen-xi-lai-yi-bo-by-jobhunter-4/
func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1.0
	} else {
		return 0.5
	}
}

func main() {
	fmt.Println(nthPersonGetsNthSeat(1))
	fmt.Println(nthPersonGetsNthSeat(2))
	fmt.Println(nthPersonGetsNthSeat(3))
}
