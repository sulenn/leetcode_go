package main

import "strings"

//https://leetcode-cn.com/problems/defanging-an-ip-address/

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
