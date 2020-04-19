package main

//https://leetcode-cn.com/problems/count-the-repetitions/

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	idx, cnt1, cnt2 := 0, 0, 0 // s2指针,s1计数,s2计数
	for cnt1 < n1 {            // 当s1的数量没有超过总个数，就可以继续读取s1
		for i := 0; i < len(s1); i++ { //遍历s1中每一个字符
			if s1[i] == s2[idx] { //如果相等指针后移
				idx++ // 指针后移
				if idx == len(s2) { //如果到最后一个
					cnt2++  //s2计数加一
					idx = 0 //指针指回开头
				}
			}
		}
		cnt1++ // 用了一个s1，计数
	}
	return cnt2 / n2
}

