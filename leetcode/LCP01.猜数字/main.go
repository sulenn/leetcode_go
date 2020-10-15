package main

//https://leetcode-cn.com/problems/guess-numbers/

func game(guess []int, answer []int) int {
	r := 0
	for i := range guess {
		if guess[i] == answer[i] {
			r++
		}
	}
	return r
}
