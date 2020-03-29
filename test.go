package main
import (
	"fmt"
	"sync"
)
type MyMutex struct {
	count int
	sync.Mutex
}
func main() {
	var mu MyMutex
	mu.Lock()

	mu.count++
	var mu2 = mu
	mu.Unlock()
	//mu2.Lock()
	mu2.count++
	mu2.Unlock()
	fmt.Println(mu.count, mu2.count)
}