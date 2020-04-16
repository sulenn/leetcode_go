package main

import (
	"fmt"
	"testing"
)

func Test_solveNQueens(T *testing.T) {
	fmt.Println(solveNQueens(1))
	fmt.Println(solveNQueens(2))
	fmt.Println(solveNQueens(3))
	fmt.Println(solveNQueens(4))
}
