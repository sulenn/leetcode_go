package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		flights := make([][]int, n)
		for v := 0; v < n; v++ {
			flights[v] = make([]int, n)
		}
		var i int
		var j int
		var price int
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		str := input.Text()
		strList := strings.Split(str, " ")
		for v := 0; v < len(strList); i++ {
			list := strings.Split(strList[i], ",")
			i, _ = strconv.Atoi(list[0])
			j, _ = strconv.Atoi(list[1])
			price, _ = strconv.Atoi(list[2])
			flights[i][j] = price
		}
		var src int
		var dst int
		var k int
		fmt.Scan(&src, &dst, &k)
		fmt.Println(findCheapestPrice(n, flights, src, dst, k))
	}
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	infCost := 2 * n * 10000
	cost := make([]int, n)
	for i := range cost {
		cost[i] = infCost
	}
	for i := 0; i <= K; i++ {
		tempCost := append(cost[:0:0], cost...)
		for _, route := range flights {
			tempCost[route[1]] = min(tempCost[route[1]], cost[route[0]]+route[2])
		}
		cost = append(tempCost[:0:0], tempCost...)
	}
	if cost[dst] >= infCost {
		return -1
	} else {
		return cost[dst]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//3
//0,1,100 1,2,100 0,2,500
//0 2 1
