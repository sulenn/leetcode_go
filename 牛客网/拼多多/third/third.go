package main

import (
	"fmt"
	"io"
	"math"
	"strings"
)

func main() {
	var N int
	var K int
	var str string
	for {
		_, err := fmt.Scan(&N,&K,&str)
		if err == io.EOF {break}
		//fmt.Println(N,K,str)
		//str = "787585"
		arr := make([]int,10)
		for i:=0; i<len(str); i++ {
			arr[str[i]-'0']++
		}
		v1, v2 := solution1(arr, K, str)
		fmt.Println(v1)
		fmt.Println(v2)
	}
}

func solution1(arr []int, K int, str string) (int, string) {
	min := math.MaxInt64
	index := 0
	for i:=0; i<len(arr); i++ {
		value := calDamage(arr, i, K)
		if  value < min {
			min = value
			index = i
		}
	}
	if arr[index] >= K {return 0, str}
	newArr := []rune(str)
	K -= arr[index]
	h := index - 1
	t := index + 1
	for K > 0 {
		if h >= 0 {
			if K >= arr[h] {
				K -= arr[h]
				for i:=len(newArr)-1; i>=0; i-- {
					if newArr[i] == rune(h+'0') {
						newArr[i] = rune(index+'0')
					}
				}
			} else {
				for i:=len(newArr)-1; i>=0 && K > 0; i-- {
					if newArr[i] == rune(h+'0') {
						newArr[i] = rune(index+'0')
						K--
					}
				}
			}
		}
		if K > 0 && t < len(arr) {
			if K >= arr[t] {
				K -= arr[t]
				for i:=0; i<len(newArr); i++ {
					if newArr[i] == rune(t+'0') {
						newArr[i] = rune(index+'0')
					}
				}
			} else {
				for i:=0; i<len(newArr) && K > 0; i++ {
					if newArr[i] == rune(t+'0') {
						newArr[i] = rune(index+'0')
						K--
					}
				}
			}
		}
		h--
		t++
	}
	return min, string(newArr)
}

func calDamage(arr []int, i int, K int) int {
	if arr[i] >= K {return 0}
	value := 0
	K -= arr[i]
	h := i-1
	t := i+1
	for K != 0 {
		if h >= 0 {
			if K>=arr[h] {
				value += (i-h)*arr[h]
				K -= arr[h]
			} else {
				value += (i-h)*K
				K = 0
			}
		}
		if K > 0 && t < len(arr) {
			if K>=arr[t] {
				value += (t-i)*arr[t]
				K -= arr[t]
			} else {
				value += (t-i)*K
				K = 0
			}
		}
		h--
		t++
	}
	strings.Re()
	return value
}

