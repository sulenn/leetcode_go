package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//华为2016校园招聘上机笔试题

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		card1 := strings.Split(strings.Split(line, "-")[0], " ")  // 牌1
		card2 := strings.Split(strings.Split(line, "-")[1], " ")  // 牌2
		if len(card1) == len(card2) {  // 同样的牌
			fmt.Println(strings.Join(card(card1, card2), " "))
			continue
		}
		if len(card1) == 2 {
			isJoker1 := judgeJoker(card1)
			if isJoker1 {  // 有王炸
				fmt.Println(strings.Join(card1, " "))
				continue
			}
		}
		if len(card2) == 2 {
			isJoker2 := judgeJoker(card2)
			if isJoker2 {  // 有王炸
				fmt.Println(strings.Join(card2, " "))
				continue
			}
		}
		if len(card1) == 4 {
			isFourSame1 := judgeFour(card1)
			if isFourSame1 {  // 有炸
				fmt.Println(strings.Join(card1, " "))
				continue
			}
		}
		if len(card2) == 4 {
			isFourSame2 := judgeFour(card2)
			if isFourSame2 {  // 有炸
				fmt.Println(strings.Join(card2, " "))
				continue
			}
		}
		fmt.Println("ERROR")
	}
}

func oneCard(card1 string, card2 string) string { // 比较牌的大小
	cardMap := map[string]int {"3":0, "4":1, "5":2, "6":3, "7":4, "8":5, "9":6, "10":7,
		"J":8,"Q":9,"K":10,"A":11, "2":12, "joker":13, "JOKER":14}
	if cardMap[card1] > cardMap[card2] {
		return card1
	}
	return card2
}

func card(card1 []string, card2 []string) []string { // 两种都是两张牌、三张牌、四张牌、五张牌
	bigger := oneCard(card1[0],card2[0])
	if bigger == card1[0] {
		return card1
	}
	return card2
}

func judgeFour(card []string) bool {  // 判断四个炸弹
	for i:=0; i<len(card)-1;i++ {
		if card[i] != card[i+1] {return false}
	}
	return true
}

func judgeJoker(card []string) bool {
	if (card[0] == "joker" && card[1] == "JOKER") || (card[0] == "JOKER" && card[1] == "joker") {
		return true
	}
	return false
}

