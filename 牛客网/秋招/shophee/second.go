package main

import (
	"fmt"
	"io"
	"sort"
)

type ch struct {
	c byte
	n int
}

func main() {
	var str string
	for {
		_, err := fmt.Scan(&str)
		if err == io.EOF {
			break
		}
		if str == "" {
			fmt.Println("")
			continue
		}
		tmap := make(map[byte]int)
		for k, _ := range str {
			tmap[str[k]]++
		}
		slice := make([]*ch, 0)
		for k, v := range tmap {
			node := new(ch)
			node.n = v
			node.c = k
			slice = append(slice, node)
		}
		sort.Slice(slice, func(i, j int) bool {
			if slice[i].c < slice[j].c {
				return true
			}
			return false
		})
		sort.Slice(slice, func(i, j int) bool {
			if slice[i].n < slice[j].n {
				return true
			}
			return false
		})
		retset := ""
		for i := 0; i < len(slice); i++ {
			for j := 0; j < slice[i].n; j++ {
				tmp := string(slice[i].c)
				retset = retset + tmp
			}
		}
		fmt.Println(retset)
	}
}
