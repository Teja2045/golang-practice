package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

// Method with pointer receiver to increment the counter
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// Method with value receiver to get the current count
func (c Counter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	counter := Counter{}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Final count:", counter.GetCount()) // Output: Final count: 1000
}
