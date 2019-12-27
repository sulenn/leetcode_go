// A concurrent prime sieve

package main

import (
	"fmt"
	"os"
)

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

func recursive(n int) int {
	if n  == 1 {
		return 1
	}
	return recursive(n-1) + n
}

func main() {
	file, err := os.Open("/null")


	if err != nil {
		fmt.Println("open error! ", err)
		return
	}
	defer func() {
		fmt.Println("close error: ")
		err := file.Close()

		if err != nil {
			fmt.Println("close error: ", err)
		} else {
			fmt.Println("close no error")
		}
	}()
}