package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出
		//c := time.Tick(time.Second * 1)
		//for range c {
		//	fmt.Println(1)
		//	defer func() {
		//		if err := recover(); err != nil {
		//			fmt.Println(err)
		//		}
		//	}()
		//	proc()
		//	fmt.Println(2)
		//}
		c := time.Tick(1 * time.Second)
		for range c {
			fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
			fmt.Println(1)
				defer func() {
					if err := recover(); err != nil {
						fmt.Println(err)}}()
				proc()
			fmt.Println(2)
		}
	}()

	select {}
}

func proc() {
	panic("ok")
}