package main

import (
	"fmt"
	"testing"
)

func Test_gameOfLife(T *testing.T) {
	temp := [][]int {{0,1,0},{0,0,1},{1,1,1},{0,0,0}}
	gameOfLife(temp)
	fmt.Println(temp)
	temp = [][]int {{0,0,0},{0,0,0},{0,0,0},{0,0,0}}
	gameOfLife(temp)
	fmt.Println(temp)
}
