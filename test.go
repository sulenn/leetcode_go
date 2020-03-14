package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(4)
	play := make(chan int)
	for i:=0; i<4; i++ {
		go run(play)
	}
	play <- 0
	wg.Wait()

}

func run(player chan int)  {
	defer wg.Done()
	num := <- player
	fmt.Println("运动员 ", num+1, " 开始跑")
	if num == 3 {
		close(player)
		return
	}
	player <- num+1
}