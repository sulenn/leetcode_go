package main

import (
	"sort"
)

//https://leetcode-cn.com/problems/sort-characters-by-frequency/description/

type bucket struct {
	word byte
	count int
}

//思路：先统计频率、然后排序
func frequencySort(s string) string {
	dic := make(map[byte]int)
	for i:=0; i<len(s);i ++ {
		dic[s[i]]++
	}
	buckets := make([]bucket, len(dic))  // 切片如果可以确定长度，最好确定长度。节约时间
	index :=0
	for k,v := range dic {
		buckets[index] = bucket{k,v}
		index++
	}
	sort.Slice(buckets, func(i, j int) bool {
		return buckets[i].count > buckets[j].count
	})
	result := []byte {}  // 字节切片的形式，优于每个字符转字符串相加
	for i:=0; i<len(buckets); i++ {
		for j:=0; j<buckets[i].count; j++ {
			result = append(result, buckets[i].word)
		}
	}
	return string(result)
	//result := ""
	//for i:=0; i<len(buckets); i++ {
	//	for j:=0; j<buckets[i].count; j++ {
	//		result += string(buckets[i].word)
	//	}
	//}
	//return result
}

//func main() {
//	arr := make([][]int, 10,10)
//	fmt.Println(arr)
//	arr1 := make([][]int, 10)
//	fmt.Println(arr1)
//}