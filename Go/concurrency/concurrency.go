package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	workItems := 5
	wg.Add(workItems)

	for i := 1; i <= workItems; i++ {
		go func(j int) {
			defer wg.Done()
			printLater(fmt.Sprintf("Hello from %d\n", j), time.Millisecond*100)
		}(i)
	}

	wg.Wait()
}

func printLater(msg string, duration time.Duration) {
	time.Sleep(duration)
	fmt.Printf(msg)
}
