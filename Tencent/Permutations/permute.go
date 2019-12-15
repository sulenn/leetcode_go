package main

func permute(nums []int) [][]int {
	numLen := len(nums)
	var result [][]int
	for i:=0 ; i<numLen; i++ {
		firstValue := []int {i}
		result = append(result, firstValue)
	}
	for j:=1; j < numLen; j++ {
		for z:=0; z < numLen; z++ {

		}
	}
}