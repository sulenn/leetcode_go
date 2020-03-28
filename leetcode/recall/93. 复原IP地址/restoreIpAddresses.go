package main

import "strings"

//https://leetcode-cn.com/problems/restore-ip-addresses/

func restoreIpAddresses(s string) []string {
	if len(s) > 12 || len(s) < 4 {return []string {}}
	result := []string {}
	backtracking(s, &result, []string {})   // 注意这儿传指针
	return result
}

func backtracking(s string, combinations *[]string, ip []string)  {
	if len(s) > (4-len(ip))*3 || len(s) < 4-len(ip) {return}  // 剩余字符串不满足要求
	if len(ip) != 0 && len(ip[len(ip)-1]) == 3 && ip[len(ip)-1] > "255" {return}   // 某一段Ip值大于255
	if len(ip) != 0 && len(ip[len(ip)-1]) > 1 && ip[len(ip)-1][0] == '0' {return}  // 某一段IP值为0开头，01、02
	if len(s) == 0 {  // 满足条件
		*combinations = append(*combinations, strings.Join(ip,"."))
		return
	}
	for i:=0; i<len(s)&&i<3; i++ {   // 最多三位，最少一位
		ip = append(ip, s[:i+1])
		backtracking(s[i+1:], combinations, ip)
		ip = ip[:len(ip)-1]
	}
}


