package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type info struct {
	s string
	good string  // 靓号，且最有价值的子字符串
	flag bool // 豹子还是顺子，豹子为true
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		arr := strings.Split(str, ",")
		infoArr := []info {}
		for _, v:= range arr {
			s := judge(v)
			if s != "" {
				if s[0] == s[1] {
					infoArr = append(infoArr,info{v, s,true})
				} else {
					infoArr = append(infoArr,info{v, s,false})
				}
			}
		}
		sort(infoArr)
		result := ""
		for _, v:= range infoArr {
			result += v.s +","
		}
		fmt.Println(strings.TrimRight(result,","))
	}
	//fmt.Println(judge("15112347234"))
}

func sort(in []info) {
	for i:=0; i<len(in); i++ {
		for j := len(in)-1; j>i; j-- {
			if !judgeNum(in[j-1], in[j]) {  // 价值最大的往上升
				in[j-1],in[j] = in[j],in[j-1]
			}
		}
	}
}

func judgeNum(info1 info, info2 info) bool {
	if len(info1.good) > len(info2.good) {
		return true
	}
	if len(info1.good) == len(info2.good) && info1.flag == true {
		return true
	}
	if len(info1.good) == len(info2.good) && info2.flag == false {
		return true
	}
	return false
}

func judge(s string) string { // 判断一个号码是否是靓号，并返回豹子或顺子
	s = s[3:]
	t := []string {}
	n := []string {}
	t_s := 0 // 豹子起始位置
	t_e := 0 // 豹子结束位置
	n_s := 0 // 顺子起始位置
	n_e := 0 // 顺子结束位子
	for i:=1; i<len(s); i++ {  // 获取豹子和顺子
		if s[i] == s[t_e] {
			t_e = i
		} else if t_e - t_s > 1 {
			t = append(t, s[t_s:t_e+1])
			t_s, t_e = i, i
		} else {
			t_s, t_e = i, i
		}
		if s[i] == s[n_e] + 1{
			n_e = i
		} else if n_e - n_s > 1 {
			n = append(n, s[n_s:n_e+1])
			n_s, n_e = i, i
		} else {
			n_s, n_e = i, i
		}
	}
	if len(t) == 0 && len(n) == 0 {  // 既无豹子也无顺子
		return ""
	}
	if len(t) == 0 && len(n) == 1 {  //只有顺子
		return n[0]
	}
	if len(t) == 0 && len(n) == 2 {  // 有两个顺子
		if len(n[0]) > len(n[1]) {  // 第一个顺子比第二顺子长
			return n[0]
		} else if len(n[0]) < len(n[1]) {  // 第二个顺子比第一个顺子长
			return n[1]
		} else {  // 一样长
			return n[1]
		}
	}
	if len(t) == 1 && len(n) == 0{  // 只有豹子
		return t[0]
	}
	if len(t) == 2 && len(n) == 0{  // 有两个豹子
		if len(t[0]) > len(t[1]) {
			return t[0]
		} else if len(t[0]) < len(t[1]) {
			return t[1]
		} else {
			return t[1]
		}
	}
	if len(t) == len(n) {  // 一个豹子一个顺子
		if len(n[0]) > len(t[0]) {
			return n[0]
		} else {
			return t[0]
		}
	}
	if len(n)  == len(t) + 1 {  //两个顺子一个豹子
		return t[0]
	}
	if len(t) == len(n)+1 {  // 两个豹子一个顺子
		if len(n[0]) == 4 {return n[0]}
		if len(t[0]) > len(t[1]) {
			return t[0]
		} else if len(t[0]) < len(t[1]) {
			return t[1]
		} else {
			return t[1]
		}
	}
	return ""
}