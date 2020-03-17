package main

//https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters/

func countCharacters(words []string, chars string) int {
	result := 0
	charDic := make(map[int32]int)
	for _, v := range chars {   // 统计 chars 中字母和对应数量
		if _,ok := charDic[v]; ok {
			charDic[v]++
		} else {
				charDic[v] = 1
		}
	}
	for _, word := range words {
		tempDic := make(map[int32]int)
		for _,v := range word {  // 统计单个单词中字母和对应数量
			if _, ok:= tempDic[v]; ok {
				tempDic[v]++
			} else {
					tempDic[v] = 1
			}
		}
		flag := true
		for k, v:= range tempDic {
			if value ,ok:= charDic[k]; !ok || v > value {  // 单词中存在字符表中没有的单词或数量不足的单词
				flag = false
			}
		}
		if flag {  // 统计满足条件的单词总长度
			result += len(word)
		}
	}
	return result
}
