package main

import "fmt"

//https://leetcode-cn.com/problems/find-common-characters/

func commonChars(A []string) []string {
	baseDic := make(map[string]int)
	for i:=0; i<len(A[0]); i++ {
		baseDic[A[0][i:i+1]]++
	}
	for i:=1; i<len(A); i++ {
		tempDic := make(map[string]int)
		for j:=0; j<len(A[i]); j++ {
			tempDic[A[i][j:j+1]]++
		}
		for k, v := range baseDic{
			tempV, ok := tempDic[k]
			if ok && v > tempV {
				baseDic[k] = tempV
			}
			if !ok {
				delete(baseDic, k)
			}
		}
	}
	result := make([]string, 0)
	for k,v := range baseDic{
		for i:=0; i<v; i++ {
			result = append(result, k)
		}
	}
	return result
}

func main() {
	fmt.Println(string(98))
}