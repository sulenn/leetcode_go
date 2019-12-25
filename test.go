// A concurrent prime sieve

package main

import "fmt"

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
	count := 0
	for i := 1; i <= 10; i++ {
		count += i
	}
	fmt.Println(count)
}