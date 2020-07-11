package 使用_goroutine_和_channel_打印五个随机数字

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

//https://github.com/lifei6671/interview-go/blob/master/question/q009.md

func main() {
	c := make(chan int)
	wg.Add(2)
	go genetateRandInt(c)
	go printInt(c)
	defer close(c)
	wg.Wait()
}

func genetateRandInt(c chan int) {
	for i := 0; i < 5; i++ {
		c <- rand.Intn(10)
	}
	wg.Done()
}

func printInt(c chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	wg.Done()
}
