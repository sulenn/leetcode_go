package main

import (
	"errors"
	"fmt"
	"strings"
)

//https://leetcode-cn.com/problems/replace-words/

//思路。hashtable 保存所有词根。计算所有词根中最长和最短词根的长度。遍历句子中的单词，以最长和最短词根为边界匹配词根。
//满足条件就替换
//func replaceWords(dict []string, sentence string) string {
//	sentence = strings.Trim(sentence," ")
//	if len(sentence) == 0 {   // 异常
//		panic(errors.New("输入的 sentence 为空句子或全空格组成的句子！"))
//	}
//	sentenceArr := strings.Split(sentence, " ")
//	dictMap := make(map[string]bool)
//	shortestWord := math.MaxInt64   // 长度最短的词根
//	longestWord := 0   // 长度最长的词根
//	for _, word := range dict {  // 构建词根的 hashtable
//		dictMap[word] = true
//		if len(word) < shortestWord {
//			shortestWord = len(word)
//		}
//		if len(word) > longestWord {
//			longestWord = len(word)
//		}
//	}
//	for i, word := range sentenceArr {  // 遍历句子中的单词
//		length := len(word)
//		for j:=shortestWord; j<=longestWord && length>=j; j++ {  // 注意判断条件中加入该单词是否越界
//			if _, ok := dictMap[word[:j]]; ok {
//				sentenceArr[i] = word[:j]
//				break
//			}
//		}
//	}
// 	return strings.Join(sentenceArr, " ")
//}

//前缀树，如何设置数中节点的数量是二维数据构建前缀树的难点
func replaceWords(dict []string, sentence string) string {
	sentence = strings.Trim(sentence," ")
	if len(sentence) == 0 {
		panic(errors.New("输入的 sentence 为空句子或全空格组成的句子！"))
	}
	sentenceArr := strings.Split(sentence, " ")

	trie := make([][]int, 10000000)   // 树中节点的数量
	//trie := make([][]int, len(dict)*26)   // 树中节点的数量
	for i, _ := range trie {
		trie[i] = make([]int,26)  //0-25对应 a-z
	}
	//color := make([]bool, len(dict)*26)   // 节点着色，如果节点为 true 说明从根节点到该节点的单词串起来是一个词根
	color := make([]bool, len(dict)*10000000)   // 节点着色，如果节点为 true 说明从根节点到该节点的单词串起来是一个词根
	countNode := 1  //下一个节点的编号，默认从 0 开始
	for _, word := range dict {   // 构建字典前缀树
		if len(word) == 0 {  // 异常
			panic(errors.New("词根中存在空串！"))
		}
		node := 0 // 根节点
		for i:=0; i<len(word); i++ {
			if trie[node][word[i]-'a'] == 0 {
				trie[node][word[i]-'a'] = countNode
				node = countNode
				countNode++
			} else {
				node = trie[node][word[i]-'a']
			}
		}
		color[node] = true
	}

	for j, word:=range sentenceArr {   // 按词根匹配替换句子中的单词
		node := 0 // 根节点
		for i:=0; i<len(word); i++ {
			if trie[node][word[i]-'a'] != 0 {
				node = trie[node][word[i]-'a' ]
				if color[node] {  // 匹配词根
					sentenceArr[j] = word[:i+1]
					break
				}
			} else {break}
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
