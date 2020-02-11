package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

//https://leetcode-cn.com/problems/replace-words/

//思路。hashtable 保存所有词根。计算所有词根中最长和最短词根的长度。遍历句子中的单词，以最长和最短词根为边界匹配词根。
//满足条件就替换
func replaceWords(dict []string, sentence string) string {
	sentence = strings.Trim(sentence," ")
	if len(sentence) == 0 {   // 异常
		panic(errors.New("输入的 sentence 为空句子或全空格组成的句子！"))
	}
	sentenceArr := strings.Split(sentence, " ")
	dictMap := make(map[string]bool)
	shortestWord := math.MaxInt64   // 长度最短的词根
	longestWord := 0   // 长度最长的词根
	for _, word := range dict {  // 构建词根的 hashtable
		dictMap[word] = true
		if len(word) < shortestWord {
			shortestWord = len(word)
		}
		if len(word) > longestWord {
			longestWord = len(word)
		}
	}
	for i, word := range sentenceArr {  // 遍历句子中的单词
		length := len(word)
		for j:=shortestWord; j<=longestWord && length>=j; j++ {  // 注意判断条件中加入该单词是否越界
			if _, ok := dictMap[word[:j]]; ok {
				sentenceArr[i] = word[:j]
				break
			}
		}
	}
 	return strings.Join(sentenceArr, " ")
}

//func join(sentenceArr []string) string {
//
//	for _, str := range sentenceArr {
//		if len(str) != 0 {}
//	}
//}


func main() {
	var temp string
	fmt.Println(temp)
	fmt.Println(len(temp))
}
