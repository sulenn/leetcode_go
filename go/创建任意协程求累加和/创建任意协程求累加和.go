package 创建任意协程求累加和

import (
"fmt"
)

func conAdd(arr []int, num int)  {  // 其中 num 是协程的数量
	ch := make(chan int, 1)
	sum := 0
	//var syn = &sync.WaitGroup{}
	//syn.Add(num)
	for i:=0; i<num; i++ {
		if i == num - 1 {
			//go add(syn, arr[len(arr)/num*i:], ch)   // 非整除，最后一次需要包括数组最后一个元素
			go addA(arr[len(arr)/num*i:], ch)   // 非整除，最后一次需要包括数组最后一个元素
		} else {
			//go add(syn, arr[len(arr)/num*i:len(arr)/num*(i+1)], ch)
			go addA(arr[len(arr)/num*i:len(arr)/num*(i+1)], ch)
		}
	}
	for i:=0; i<num; i++ {
		sum += <- ch
	}
	//syn.Wait()
	fmt.Println(sum)
}

//func add(syn *sync.WaitGroup, arr []int, ch chan int)  {
func addA(arr []int, ch chan int)  {
	sum := 0
	for _, v:= range arr {
		sum += v
	}
	ch <- sum
	//syn.Done()
}

func main()  {
	conAdd([]int {1,2,3,4,5,6,7,8,9}, 2)
	conAdd([]int {1,2,3,4,5,6,7,8,9}, 3)
	conAdd([]int {1,2,3,4,5,6,7,8,9}, 4)
	conAdd([]int {1,2,3,4,5,6,7,8,9}, 5)
	conAdd([]int {1,2,3,4,5,6,7,8,9}, 6)
	conAdd([]int {1,2,3,4,5,6,7,8,9}, 7)
	conAdd([]int {1,2,3,4,5,6,7,8,9}, 8)
	conAdd([]int {1,2,3,4,5,6,7,8,9}, 9)
	conAdd([]int {1,2,3,4,5,6,7,8,9}, 10)
}
