package main

import (
	"fmt"
	"sync"
)

func produceRand(wg *sync.WaitGroup,ch chan int)  {
	defer wg.Done()
	for i:=0; i<5 ; i++ {
		ch <- i
	}
	close(ch)
}

func printChan(wg *sync.WaitGroup, ch chan int)  {
	defer wg.Done()
	for v := range ch{
		fmt.Println(v)
	}
}

func main()  {
	ch := make(chan int, 1)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go produceRand(wg, ch)
	go printChan(wg, ch)
	wg.Wait()
}