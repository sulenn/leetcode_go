package main

import (
	"fmt"
	"testing"
)

func Test_pacificAtlantic(T *testing.T) {
	//fmt.Println(pacificAtlantic([][]int {{3,1},{1,2}}))
	//fmt.Println(pacificAtlantic([][]int {{1,2,2,3,5},{3,2,3,2,4},{2,4,5,3,1},{6,7,1,4,5},{5,1,1,2,4}}))
	fmt.Println(pacificAtlantic([][]int{{2, 3, 5}, {3, 2, 4}, {5, 3, 1}}))
}
