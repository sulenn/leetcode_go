package main

import "fmt"

func romanToInt(s string) int {
	var result = 0;
	//var dic = make(map[byte]int);   //当这里的键为 byte 类型时，下面取字符串中字符直接用下标。如果键为 string，则需要用切片类型
	//dic['I'] = 1;
	//dic['V'] = 5;
	//dic['X'] = 10;
	//dic['L'] = 50;
	//dic['C'] = 100;
	//dic['D'] = 500;
	//dic['M'] = 1000;
	var dic = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := 0; i < len(s); i++ {
		if (i == len(s) - 1) {
			result += dic[s[i]]
		} else {
			if (dic[s[i]] < dic[s[i+1]]) {
				result -= dic[s[i]]
			} else {
				result += dic[s[i]]
			}
		}
	}
	return result;
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
